package num2words

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	translate "github.com/Humpy-JBRM/elevenlabs/facade/translate"
)

type TextProcessor interface {
	Process(text string) (string, string)
}

type textProcessFunc func([]string) []string

type processor struct {
	reMatch     *regexp.Regexp
	processFunc textProcessFunc
}

type textProcessor struct {
	processorsByName map[string]*processor
	targetLanguage   string
}

func NewTextProcessor(targetLanguage string) TextProcessor {
	p := &textProcessor{
		targetLanguage: targetLanguage,
		// TODO: deal with decimal numbers
		// TODO: dates (UK format + US format)   01/03/2023 == 1 March (UK) and Jan 3 (US)
		// TODO: years, e.g. "2023" is "twenty twenty three" and not "two thousand twenty three"
		//       this needs to be figured out from the context, so you'll need a clever bit of
		//       sentence analysis
	}
	p.processorsByName = map[string]*processor{
		"spell_numbers":     &processor{reMatch: regexp.MustCompile("([0-9]+)(.*)"), processFunc: p.preprocessNumber},
		"exclamation_marks": &processor{reMatch: regexp.MustCompile("(.*)\\!$"), processFunc: p.preprocessExcalamation},
		"domain_names":      &processor{reMatch: regexp.MustCompile("(.*\\..*)\\."), processFunc: p.preprocessDomain},
	}
	return p
}

func (c *textProcessor) Process(original string) (string, string) {
	// Step 1: Split the text into words.
	// This assumes that:
	//
	//	- the text is in a latinised language
	//	- "words" (tokens) are separated by spaces
	result := make([]string, 0)
	for _, word := range strings.Split(strings.TrimSpace(original), " ") {
		// Trim all space and strip commas: " 1,234 " => "1234"
		trimmed := strings.TrimSpace(word)
		commasStripped := strings.ReplaceAll(trimmed, ",", "")
		matched := false
		for _, preProcessor := range c.processorsByName {
			matches := preProcessor.reMatch.FindAllStringSubmatch(commasStripped, -1)
			if len(matches) > 0 {
				matched = true
				processed := preProcessor.processFunc(matches[0])
				result = append(result, processed...)
				break
			}
		}
		if !matched {
			result = append(result, trimmed)
		}
	}
	return original, strings.Join(result, " ")
}

func (c *textProcessor) preprocessNumber(text []string) []string {
	intValue, err := strconv.Atoi(text[0])
	if err != nil {
		// Somehow this is not a parseable number.
		// Most likely because it is a floating point number, e.g. 3.1415926
		return text
	}
	asWords := Convert(intValue)
	if c.targetLanguage != "" {
		translator, err := translate.NewTranslateProcessorFactory().SourceLanguage("en").TargetLanguage(c.targetLanguage).New()
		if err != nil {
			log.Println(err.Error())
			return text
		}
		asWords, err = translator.Execute(asWords)
		if err != nil {
			log.Println(err.Error())
			return text
		}
	}
	return strings.Split(asWords, " ")
}

func (c *textProcessor) preprocessDomain(text []string) []string {
	return []string{text[1] + " ."}
}

func (c *textProcessor) preprocessExcalamation(text []string) []string {
	return []string{text[1] + " !"}
}

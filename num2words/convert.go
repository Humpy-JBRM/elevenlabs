package num2words

import (
	"regexp"
	"strconv"
	"strings"
)

type NumberConverter interface {
	Convert(text string) (string, string)
}

type numberConverter struct {
	reInteger *regexp.Regexp
}

func NewNumberConverter() NumberConverter {
	return &numberConverter{
		// TODO: deal with decimal numbers
		// TODO: dates (UK format + US format)   01/03/2023 == 1 March (UK) and Jan 3 (US)
		// TODO: years, e.g. "2023" is "twenty twenty three" and not "two thousand twenty three"
		//       this needs to be figured out from the context, so you'll need a clever bit of
		//       sentence analysis
		reInteger: regexp.MustCompile("([0-9]+)(.*)"),
	}
}

func (c *numberConverter) Convert(original string) (string, string) {
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
		matches := c.reInteger.FindAllStringSubmatch(commasStripped, -1)
		if len(matches) > 0 {
			intValue, err := strconv.Atoi(commasStripped)
			if err != nil {
				// Somehow this is not a parseable number.
				// Most likely because it is a floating point number, e.g. 3.1415926
				result = append(result, trimmed)
				continue
			}

			for _, w := range strings.Split(Convert(intValue), " ") {
				result = append(result, w)
			}
			continue
		}

		result = append(result, trimmed)
	}
	return original, strings.Join(result, " ")
}

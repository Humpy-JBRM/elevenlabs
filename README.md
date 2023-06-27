# elevenlabs
## Convert Arabic Numerals to Words (with optional translate)

When generating audio in a language other than English, Elevenlabs says all numbers in english.

For example, `me comi 50 ratones` (I ate 50 mice) would be generated as `me comi *FIFTY* ratones` when what you actually want is `me comi cincuenta ratones`

(see https://help.elevenlabs.io/hc/en-us/articles/14888917355409-Why-are-some-numbers-and-words-not-properly-pronounced-in-the-correct-language-)

The recommended fix is to spell out all numbers - i.e. replace them with the number in words:

	- `I ate 50 mice`

becomes

	- `I ate fifty mice`

then run that through whatever translation you're using before sending it to elevenlabs.

This golang client makes it nice and easy:

	- converts numbers to their word form

	- translates the resulting sentence into other languages (needs GOOGLE_APPLICATION_CREDENTIALS)

	- makes the call to the elevenlabs API (needs a elevenlabs API key)

I wrote this to solve a problem that I had and thought it might be useful to others.

## Run from the command line:

### Pre-requisites:

	1: XI_API_KEY environment variable == value of your api key
		(https://docs.elevenlabs.io/api-reference/quick-start/authentication)

	2: GOOGLE_APPLICATION_CREDENTIALS environment variable == path of credentials.json for an account with translate API enabled

	$ go run main.go I saw 5,100 mice
	IN : I saw 5,100 mice
	OUT: I saw five thousand one hundred mice
	Audio written to out.mp3

## Example code
```go

import (
	"fmt"
	"io"
	"log"
	translate "github.com/Humpy-JBRM/elevenlabs/facade/translate"
	tts "github.com/Humpy-JBRM/elevenlabs/facade/tts"
	"github.com/Humpy-JBRM/elevenlabs/num2words"
	"path/filepath"
	"strings"
)

func main() {
	// convert all arabic numbers to words
	sentence := "I saw 5,100 mice"
	original, converted := num2words.NewNumberConverter().Convert(sentence)

	// translate the converted sentence
	// let's do spanish, just for fun
	targetLanguage := "es"
	absDot, _ := filepath.Abs(".")
	// GOOGLE_APPLICATION_CREDENTIALS must be set
	translator, err := translate.NewTranslateProcessorFactory().SourceLanguage("en").TargetLanguage(targetLanguage).New()
	if err != nil {
		log.Fatal(err)
	}
	translated, err := translator.Execute(converted)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("IN : %s\n", original)
	fmt.Printf("OUT: %s\n", translated)

	// Convert this to a voice via eleven labs
	voiceId := "21m00Tcm4TlvDq8ikWAM"
	audio, err := tts.
		NewTextToSpeech(os.Getenv("XI_API_KEY"), "eleven_monolingual_v1").
		Generate(translated, voiceId, 0)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(filepath.Join(absDot, "out.mp3"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = io.Copy(f, audio)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Audio written to %s\n", filepath.Join(absDot, "out.mp3"))
}
```

See main.go for a complete working end-to-end example



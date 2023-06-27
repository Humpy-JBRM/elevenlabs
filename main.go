package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	translate "github.com/Humpy-JBRM/elevenlabs/facade/translate"
	tts "github.com/Humpy-JBRM/elevenlabs/facade/tts"
	"github.com/Humpy-JBRM/elevenlabs/num2words"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: main sentence...")
	}

	// convert all arabic numbers to words
	sentence := strings.Join(os.Args[1:], " ")
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

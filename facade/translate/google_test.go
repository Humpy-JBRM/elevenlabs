package facade

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestTranslate(t *testing.T) {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", filepath.Join(os.Getenv("HOME"), "translate-credentials.json"))
	translator, err := NewTranslateProcessorFactory().SourceLanguage("en").TargetLanguage("es").New()
	if err != nil {
		t.Fatal(err)
	}
	sentence := "I saw 538 mice in each of my 12 kitchens"
	translated, err := translator.Execute(sentence)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("IN : %s\n", sentence)
	fmt.Printf("OUT: %s\n", translated)
}

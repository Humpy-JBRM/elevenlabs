package num2words

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSimpleSentence(t *testing.T) {
	sentence := "I saw 500 mice"
	expected := "I saw five hundred mice"
	_, actual := NewTextProcessor("").Process(sentence)
	if expected != actual {
		t.Errorf("\nExpected: '%s'\nActual  : '%s'", expected, actual)
	}
}

func TestSimpleSentenceCommaSeparated(t *testing.T) {
	sentence := "I saw 5,100 mice"
	expected := "I saw five thousand one hundred mice"
	_, actual := NewTextProcessor("").Process(sentence)
	if expected != actual {
		t.Errorf("\nExpected: '%s'\nActual  : '%s'", expected, actual)
	}
}

// This test models the use-case where the sentence has already been translated
// and has numbers
func TestSimpleSentenceTranslateNumbers(t *testing.T) {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", filepath.Join(os.Getenv("HOME"), "translate-credentials.json"))
	sentence := "vi 5100 ratones!"
	expected := "vi cinco mil cien ratones !"
	_, actual := NewTextProcessor("es").Process(sentence)
	if expected != actual {
		t.Errorf("\nExpected: '%s'\nActual  : '%s'", expected, actual)
	}
}

func TestDomainName(t *testing.T) {
	sentence := "I work at microsoft.com."
	expected := "I work at microsoft.com ."
	_, actual := NewTextProcessor("").Process(sentence)
	if expected != actual {
		t.Errorf("\nExpected: '%s'\nActual  : '%s'", expected, actual)
	}
}

func TestExclamation(t *testing.T) {
	sentence := "I work at microsoft.com!"
	expected := "I work at microsoft.com !"
	_, actual := NewTextProcessor("").Process(sentence)
	if expected != actual {
		t.Errorf("\nExpected: '%s'\nActual  : '%s'", expected, actual)
	}
}

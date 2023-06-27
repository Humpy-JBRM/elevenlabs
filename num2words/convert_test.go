package num2words

import "testing"

func TestSimpleSentence(t *testing.T) {
	sentence := "I saw 500 mice"
	expected := "I saw five hundred mice"
	_, actual := NewTextProcessor().Process(sentence)
	if expected != actual {
		t.Errorf("\nExpected: '%s'\nActual  : '%s'", expected, actual)
	}
}

func TestSimpleSentenceCommaSeparated(t *testing.T) {
	sentence := "I saw 5,100 mice"
	expected := "I saw five thousand one hundred mice"
	_, actual := NewTextProcessor().Process(sentence)
	if expected != actual {
		t.Errorf("\nExpected: '%s'\nActual  : '%s'", expected, actual)
	}
}

func TestDomainName(t *testing.T) {
	sentence := "I work at microsoft.com."
	expected := "I work at microsoft.com ."
	_, actual := NewTextProcessor().Process(sentence)
	if expected != actual {
		t.Errorf("\nExpected: '%s'\nActual  : '%s'", expected, actual)
	}
}

func TestExclamation(t *testing.T) {
	sentence := "I work at microsoft.com!"
	expected := "I work at microsoft.com !"
	_, actual := NewTextProcessor().Process(sentence)
	if expected != actual {
		t.Errorf("\nExpected: '%s'\nActual  : '%s'", expected, actual)
	}
}

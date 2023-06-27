package num2words

import "testing"

func TestSimpleSentence(t *testing.T) {
	sentence := "I saw 500 mice"
	expected := "I saw five hundred mice"
	_, actual := NewNumberConverter().Convert(sentence)
	if expected != actual {
		t.Errorf("\nExpected: '%s'\nActual  : '%s'", expected, actual)
	}
}

func TestSimpleSentenceCommaSeparated(t *testing.T) {
	sentence := "I saw 5,100 mice"
	expected := "I saw five thousand one hundred mice"
	_, actual := NewNumberConverter().Convert(sentence)
	if expected != actual {
		t.Errorf("\nExpected: '%s'\nActual  : '%s'", expected, actual)
	}
}

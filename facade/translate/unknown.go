package facade

import (
	"fmt"
)

type unknownTranslateProcessor struct {
}

func (ip *unknownTranslateProcessor) Execute(text string) (string, error) {
	return "", fmt.Errorf("Process(): unknown implementation")
}

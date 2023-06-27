package facade

import (
	"fmt"
)

type azureTranslateProcessor struct {
	sourceLanguage string
	targetLanguage string
}

func NewAzureTranslateProcessor(factory *TranslateProcessorFactoryImpl) (TranslateProcessor, error) {
	processor := &azureTranslateProcessor{
		sourceLanguage: factory.sourceLanguage,
		targetLanguage: factory.targetLanguage,
	}

	return processor, nil
}

func (gtp *azureTranslateProcessor) Execute(text string) (string, error) {
	return "", fmt.Errorf("translate.Execute(): Implement AZURE translate")
}

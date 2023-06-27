package facade

import (
	"fmt"
)

type awsTranslateProcessor struct {
	sourceLanguage string
	targetLanguage string
}

func NewAwsTranslateProcessor(factory *TranslateProcessorFactoryImpl) (TranslateProcessor, error) {
	processor := &awsTranslateProcessor{
		sourceLanguage: factory.sourceLanguage,
		targetLanguage: factory.targetLanguage,
	}

	return processor, nil
}

func (gtp *awsTranslateProcessor) Execute(text string) (string, error) {
	return "", fmt.Errorf("translate.Execute(): Implement AWS translate")
}

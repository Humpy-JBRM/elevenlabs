package facade

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

type googleTranslateProcessor struct {
	processorUrl    string
	authToken       string
	credentialsJson string
	sourceLanguage  string
	targetLanguage  string
}

func NewGoogleTranslateProcessor(factory *TranslateProcessorFactoryImpl) (TranslateProcessor, error) {
	processor := &googleTranslateProcessor{
		credentialsJson: os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"),
		sourceLanguage:  factory.sourceLanguage,
		targetLanguage:  factory.targetLanguage,
	}

	return processor, nil
}

func (gtp *googleTranslateProcessor) Execute(text string) (string, error) {
	ctx := context.Background()

	lang, err := language.Parse(gtp.targetLanguage)
	if err != nil {
		return "", fmt.Errorf("language.Parse: %w", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", fmt.Errorf("Translate: %w", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("Translate returned empty response to text: %s", text)
	}
	return resp[0].Text, nil
}

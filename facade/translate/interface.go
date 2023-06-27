package facade

type TranslateProcessorFactory interface {
	SourceLanguage(sourceLanguage string) TranslateProcessorFactory
	TargetLanguage(targetLanguage string) TranslateProcessorFactory
	New() (TranslateProcessor, error)
}

type TranslateProcessor interface {
	Execute(text string) (string, error)
}

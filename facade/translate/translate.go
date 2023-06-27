package facade

func NewTranslateProcessorFactory() TranslateProcessorFactory {
	tpf := &TranslateProcessorFactoryImpl{
		ipType: TRANSLATE_PROCESSOR_GOOGLE,
	}
	return tpf
}

type TranslateProcessorType string

const (
	TRANSLATE_PROCESSOR_GOOGLE = "google"
)

type TranslateProcessorFactoryImpl struct {
	ipType         TranslateProcessorType
	sourceLanguage string
	targetLanguage string
}

func (tpf *TranslateProcessorFactoryImpl) SourceLanguage(sourceLanguage string) TranslateProcessorFactory {
	tpf.sourceLanguage = sourceLanguage
	return tpf
}

func (tpf *TranslateProcessorFactoryImpl) TargetLanguage(targetLanguage string) TranslateProcessorFactory {
	tpf.targetLanguage = targetLanguage
	return tpf
}

func (tpf *TranslateProcessorFactoryImpl) New() (TranslateProcessor, error) {
	switch tpf.ipType {
	case TRANSLATE_PROCESSOR_GOOGLE:
		return NewGoogleTranslateProcessor(tpf)
	}

	return &unknownTranslateProcessor{}, nil
}

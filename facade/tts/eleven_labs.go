package facade

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Humpy-JBRM/elevenlabs/util"
)

type TextToSpeech interface {
	Generate(text string, voiceId string, optimiseLatency int) (io.Reader, error)
}

type textToSpeech struct {
	apiKey  string
	modelId string
}

type VoiceSettings struct {
	Stability       int `json:"stability"`
	SimilarityBoost int `json:"similarity_boost"`
}

type TTSRequest struct {
	Text          string        `json:"text"`
	ModelId       string        `json:"model_id"`
	VoiceSettings VoiceSettings `json:"voice_settings"`
}

func NewTextToSpeech(apiKey string, modelId string) TextToSpeech {
	return &textToSpeech{
		apiKey:  apiKey,
		modelId: modelId,
	}
}

func (t *textToSpeech) Generate(text string, voiceId string, optimiseLatency int) (io.Reader, error) {
	payload := TTSRequest{
		Text:    text,
		ModelId: t.modelId,
		VoiceSettings: VoiceSettings{
			Stability:       0,
			SimilarityBoost: 0,
		},
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("Generate(): %s", err.Error())
	}
	url := fmt.Sprintf("https://api.elevenlabs.io/v1/text-to-speech/%s?optimize_streaming_latency=%d", voiceId, optimiseLatency)
	code, body, err := util.HttpPost(
		url,
		payloadBytes,
		"Accept: audio/mpeg",
		"xi-api-key: "+t.apiKey,
		"Content-type: application/json",
	)
	if err != nil {
		return nil, fmt.Errorf("Generate(): %s", err.Error())
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf("Generate(): API returned status %d", code)
	}
	return bytes.NewReader(body), nil
}

package facade

import (
	"io"
	"os"
	"testing"
)

func TestElevenLabs(t *testing.T) {
	sentence := "I saw five thousand one hundred mice"
	voiceId := "21m00Tcm4TlvDq8ikWAM"
	audio, err := NewTextToSpeech(os.Getenv("XI_API_KEY"), "eleven_monolingual_v1").Generate(sentence, voiceId, 0)
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.OpenFile("out.mp3", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	io.Copy(f, audio)
}

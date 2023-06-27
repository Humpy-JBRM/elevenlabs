package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func HttpPost(theUrl string, payload []byte, header ...string) (int, []byte, error) {
	var netClient = &http.Client{
		Timeout: time.Second * 60,
	}

	// Authenticate
	var request *http.Request
	var err error
	request, err = http.NewRequest("POST", theUrl, bytes.NewBuffer(payload))
	if err != nil {
		return 0, []byte{}, err
	}

	// Add the headers
	for _, h := range header {
		fields := strings.Split(strings.TrimSpace(h), ":")
		if len(fields) != 2 {
			return 0, []byte{}, fmt.Errorf("Invalid header: %s", h)
		}

		request.Header.Add(strings.TrimSpace(fields[0]), strings.TrimSpace(fields[1]))
	}

	response, err := netClient.Do(request)
	if err != nil {
		code := 0
		if response != nil {
			code = response.StatusCode
		}
		return code, []byte{}, err
	}
	if response == nil {
		return 0, []byte{}, fmt.Errorf("ERROR: %s: Got nil response from server", theUrl)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	return response.StatusCode, responseBytes, err
}

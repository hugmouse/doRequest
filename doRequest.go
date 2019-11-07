package doRequest

import (
	"encoding/json"
	"io"
	"net/http"
)

func RequestDecode(method, url, token string, body io.Reader, decodeTo interface{}) error {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return err
	}

	if token != "" {
		req.Header.Set("Authorization", token)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&decodeTo)
}



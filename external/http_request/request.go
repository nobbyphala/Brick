package http_request

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type httpRequest struct {
}

func NewHttpRequest() *httpRequest {
	return &httpRequest{}
}

func (ht httpRequest) Post(ctx context.Context, url string, headers map[string]string, body interface{}, response interface{}) error {
	client := &http.Client{}

	jsonBody, err := json.Marshal(body)
	if nil != err {
		log.Println("error marshal request body")
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Println("Error creating request:", err)
		return err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	//additional headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("error making POST request:", err)
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("errror closing body", err)
		}
	}(resp.Body)

	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		log.Println("error decoding response body:", err)
		return err
	}

	return nil
}

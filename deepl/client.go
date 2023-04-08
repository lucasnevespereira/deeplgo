package deepl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

type Client struct {
	http    *http.Client
	baseUrl string
	authKey string
}

func NewClient(authKey string) *Client {
	client := Client{
		http:    &http.Client{},
		baseUrl: "https://api-free.deepl.com/v2",
		authKey: fmt.Sprintf("DeepL-Auth-Key %s", authKey),
	}

	return &client
}

func (c *Client) Translate(texts []string, sourceLang, targetLang string) (*TranslateResponse, error) {
	url := fmt.Sprintf("%s/translate", c.baseUrl)
	body := &TranslateRequest{
		Text:       texts,
		SourceLang: sourceLang,
		TargetLang: targetLang,
	}
	rawBody, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "marshal translate request body")
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(rawBody))
	if err != nil {
		return nil, errors.Wrap(err, "create translate POST request")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", c.authKey)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "call translate POST request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("status_unknown %d, url: %s, response: %s", resp.StatusCode, url, resp.Body)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read translate POST response body")
	}

	var translateResponse *TranslateResponse
	if err = json.Unmarshal(bodyBytes, &translateResponse); err != nil {
		return nil, errors.Wrap(err, "Unmarshal response body")
	}

	return translateResponse, err
}

package deepai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type Colorizer struct {
	apiKey string
	client *http.Client
}

func New(apiKey string, client *http.Client) Colorizer {
	if client == nil {
		client = http.DefaultClient
	}
	return Colorizer{
		apiKey: apiKey,
		client: client,
	}
}

func (c *Colorizer) Colorize(b []byte) ([]byte, error) {
	if c.apiKey == "" {
		return nil, errors.New("empty apiKey")
	}
	req, errReq := c.buildRequest(b)
	if errReq != nil {
		return nil, errors.New("can't build request")
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed post resuest to deep.ai %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("incorrect status " + resp.Status)
	}
	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read response: %w", err)
	}
	var deepResp deepAIResp
	err = json.Unmarshal(bodyResp, &deepResp)
	if err != nil {
		return nil, fmt.Errorf("can't unpack json: %w", err)
	}

	return c.fromURL(deepResp.OutputURL)
}

func (c *Colorizer) buildRequest(b []byte) (*http.Request, error) {
	byteReader := bytes.NewReader(b)
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)
	fw, err := w.CreateFormFile("image", "test.png")
	if err != nil {
		return nil, fmt.Errorf("can't read from stream %w", err)
	}
	_, err = io.Copy(fw, byteReader)
	w.Close()
	req, _ := http.NewRequest(http.MethodPost, "https://api.deepai.org/api/colorizer", buf)
	req.Header.Set("api-key", c.apiKey)
	req.Header.Set("Content-Type", w.FormDataContentType())
	return req, nil
}

func (c *Colorizer) fromURL(url string) ([]byte, error) {
	deepColorized, err := c.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("can't read from outputUrl %w url:%s", err, url)
	}
	defer deepColorized.Body.Close()
	if deepColorized.StatusCode != http.StatusOK {
		return nil, errors.New("not found image from url")
	}
	respBody, errRead := ioutil.ReadAll(deepColorized.Body)
	if errRead != nil {
		return nil, errors.New("can't read from resp")
	}
	return respBody, nil
}

type deepAIResp struct {
	ID        string `json:"id"`
	OutputURL string `json:"output_url"`
}

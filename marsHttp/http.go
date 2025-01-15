package marsHttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/marsli9945/mars-go/marsLog"
	"io"
	"net/http"
)

var client = &http.Client{}

func Get(url string) (string, error) {
	req, err := getRequest(url, nil, http.MethodGet)
	if err != nil {
		return "", err
	}
	body, err := doClient(req)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func GetAndHeaderForStruct(url string, headers map[string]string, result any) error {
	req, err := getRequest(url, nil, http.MethodGet)
	if err != nil {
		return err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	body, err := doClient(req)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, result)
	if err != nil {
		return err
	}
	return nil
}

func Post(url string, data map[string]any) (string, error) {
	req, err := getRequest(url, data, http.MethodPost)
	if err != nil {
		return "", err
	}
	body, err := doClient(req)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func PostAndHeaderForStruct(url string, data map[string]any, headers map[string]string, result any) error {
	req, err := getRequest(url, data, http.MethodPost)
	if err != nil {
		return err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	body, err := doClient(req)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, result)
	if err != nil {
		return err
	}
	return nil
}

func getRequest(url string, data any, method string) (*http.Request, error) {
	if method == http.MethodGet || data == nil {
		return http.NewRequest(method, url, nil)
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func doClient(req *http.Request) ([]byte, error) {
	response, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			marsLog.Logger().ErrorF("doClient Close Error: %v, URL: %s, Method: %s", err, req.URL, req.Method)
		}
	}(response.Body)

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return []byte{}, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

package service

import (
	"io/ioutil"
	"net/http"
)

type HTTPClient interface {
	GetData(string) ([]byte, error)
}

type HTTPClientImpl struct{}

func (hc *HTTPClientImpl) GetData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()
	return body, err
}

package api

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/k1rnt/mediawikiAPI/config"
)

func init() {
	config.Env_load()
}

func getRequest(cli *http.Client, params url.Values) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, config.Endpoint, nil)
	if err != nil {
		return nil, err
	}
	request.URL.RawQuery = params.Encode()

	response, err := cli.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func postRequest(cli *http.Client, data url.Values) ([]byte, error) {
	var b *strings.Reader = nil
	if data != nil {
		b = strings.NewReader(data.Encode())
	}
	request, err := http.NewRequest(http.MethodPost, config.Endpoint, b)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	response, err := cli.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

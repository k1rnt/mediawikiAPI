package api

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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

func postFileRequest(cli *http.Client, params map[string]string, paramName, path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	out := &bytes.Buffer{}
	writer := multipart.NewWriter(out)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, config.Endpoint, out)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Set("Content-Type", writer.FormDataContentType())
	// request.Header.Set("Content-type", "multipart/form-data; charset=UTF-8")

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

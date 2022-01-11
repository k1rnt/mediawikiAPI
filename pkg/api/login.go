package api

import (
	"net/http"
	"net/url"

	"github.com/k1rnt/mediawikiAPI/config"
)

func GetLoginRequest(cli *http.Client, token string) ([]byte, error) {
	data := url.Values{}
	data.Add("action", "login")
	data.Add("lgname", config.User)
	data.Add("lgpassword", config.Pass)
	data.Add("lgtoken", token)
	data.Add("format", "json")

	body, err := postRequest(cli, data)
	if err != nil {
		return nil, err
	}

	return body, nil
}

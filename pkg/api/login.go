package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/k1rnt/mediawikiAPI/config"
)

type LoginToken struct {
	Batchcomplete string `json:"batchcomplete"`
	Query         struct {
		Tokens struct {
			Logintoken string `json:"logintoken"`
		} `json:"tokens"`
	} `json:"query"`
}

func GetLoginToken(cli *http.Client) LoginToken {
	request, err := http.NewRequest(http.MethodGet, config.Endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	params := request.URL.Query()
	params.Add("action", "query")
	params.Add("meta", "tokens")
	params.Add("type", "login")
	params.Add("format", "json")
	request.URL.RawQuery = params.Encode()

	response, err := cli.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var logintoken LoginToken
	if err := json.Unmarshal(body, &logintoken); err != nil {
		log.Fatal(err)
	}

	return logintoken
}

func GetLoginRequest(cli *http.Client, token string) []byte {
	data := url.Values{}
	data.Add("action", "login")
	data.Add("lgname", config.User)
	data.Add("lgpassword", config.Pass)
	data.Add("lgtoken", token)
	data.Add("format", "json")

	request, err := http.NewRequest(http.MethodPost, config.Endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	response, err := cli.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

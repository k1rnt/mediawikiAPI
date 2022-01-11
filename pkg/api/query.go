package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

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

type CSRFToken struct {
	Batchcomplete string `json:"batchcomplete"`
	Query         struct {
		Tokens struct {
			Csrftoken string `json:"csrftoken"`
		} `json:"tokens"`
	} `json:"query"`
}

func CsrfToken(cli *http.Client) CSRFToken {
	request, err := http.NewRequest(http.MethodGet, config.Endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	params := request.URL.Query()
	params.Add("action", "query")
	params.Add("meta", "tokens")
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

	var csrftoken CSRFToken
	if err := json.Unmarshal(body, &csrftoken); err != nil {
		log.Fatal(err)
	}

	return csrftoken
}

package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
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
	params := url.Values{}
	params.Add("action", "query")
	params.Add("meta", "tokens")
	params.Add("type", "login")
	params.Add("format", "json")

	body, err := getRequest(cli, params)
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
	params := url.Values{}
	params.Add("action", "query")
	params.Add("meta", "tokens")
	params.Add("format", "json")

	body, err := getRequest(cli, params)
	if err != nil {
		log.Fatal(err)
	}

	var csrftoken CSRFToken
	if err := json.Unmarshal(body, &csrftoken); err != nil {
		log.Fatal(err)
	}

	return csrftoken
}

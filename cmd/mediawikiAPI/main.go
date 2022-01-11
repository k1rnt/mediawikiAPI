package main

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/k1rnt/mediawikiAPI/pkg/api"
	"github.com/k1rnt/mediawikiAPI/pkg/utils"
)

func setClient() *http.Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Jar:     jar,
		Timeout: timeout,
	}
	return client
}

func main() {
	// example#1
	cli := setClient()
	logintoken := api.GetLoginToken(cli).Query.Tokens.Logintoken
	api.GetLoginRequest(cli, logintoken)
	csrftoken := api.CsrfToken(cli).Query.Tokens.Csrftoken

	files, err := utils.FindWikiFiles("dir")
	if err != nil {
		log.Fatal(err)
	}
	for _, path := range files {
		api.EditRequest(cli, csrftoken, path)
	}
}

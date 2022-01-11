package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/k1rnt/mediawikiAPI/pkg/api"
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
	cli := setClient()
	logintoken := api.GetLoginToken(cli).Query.Tokens.Logintoken
	api.GetLoginRequest(cli, logintoken)
	fmt.Printf("%s\n", api.CsrfToken(cli).Query.Tokens.Csrftoken)
}

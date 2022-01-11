package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/k1rnt/mediawikiAPI/config"
)

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

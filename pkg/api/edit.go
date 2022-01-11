package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/k1rnt/mediawikiAPI/config"
	"github.com/k1rnt/mediawikiAPI/pkg/utils"
)

func EditRequest(cli *http.Client, csrftoken string, path string) []byte {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to open %s\n", path)
	}
	filename := utils.GetFileNameWithoutExt(path)

	data := url.Values{}
	data.Add("action", "edit")
	data.Add("title", filename)
	data.Add("text", string(f))
	data.Add("token", csrftoken)
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

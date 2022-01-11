package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/k1rnt/mediawikiAPI/pkg/utils"
	"golang.org/x/xerrors"
)

func EditRequest(cli *http.Client, csrftoken string, path string) ([]byte, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, xerrors.Errorf("Failed to open %s\n", path)
	}
	filename, err := utils.GetFileNameWithoutExt(path)
	if err != nil {
		return nil, err
	}

	data := url.Values{}
	data.Add("action", "edit")
	data.Add("title", filename)
	data.Add("text", string(f))
	data.Add("token", csrftoken)
	data.Add("format", "json")

	body, err := postRequest(cli, data)
	if err != nil {
		log.Fatal(err)
	}

	return body, nil
}

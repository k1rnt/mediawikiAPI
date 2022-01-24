package handler

import (
	"fmt"
	"net/http"

	"github.com/k1rnt/mediawikiAPI/config"
	"github.com/k1rnt/mediawikiAPI/pkg/api"
	"github.com/k1rnt/mediawikiAPI/pkg/utils"
	"golang.org/x/xerrors"
)

func UploadPageHandler(cli *http.Client, token, path string) error {
	files, err := utils.FindWikiFiles(path)
	if err != nil {
		return xerrors.Errorf("%s", err.Error())
	}
	for _, path := range files {
		if _, err := api.EditRequest(cli, token, path); err != nil {
			return xerrors.Errorf("%s", err.Error())
		}
		fmt.Printf("success %s\n", path)
		fmt.Printf("target to %s\n", config.Endpoint)
	}
	return nil
}

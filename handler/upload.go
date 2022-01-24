package handler

import (
	"fmt"
	"net/http"

	"github.com/k1rnt/mediawikiAPI/pkg/api"
	"github.com/k1rnt/mediawikiAPI/pkg/utils"
	"golang.org/x/xerrors"
)

func UploadImageHandler(cli *http.Client, token, path string) error {
	images, err := utils.FindImageFiles(path)
	if err != nil {
		return xerrors.Errorf("%s", err.Error())
	}
	for _, path := range images {
		body, err := api.Upload(cli, token, path)
		if err != nil {
			return xerrors.Errorf("%s", err.Error())
		}
		fmt.Printf("%s\n", body)
		fmt.Printf("success %s\n", path)
	}
	return nil
}

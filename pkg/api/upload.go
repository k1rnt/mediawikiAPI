package api

import (
	"log"
	"net/http"
	"path/filepath"
)

func Upload(cli *http.Client, csrftoken, path string) ([]byte, error) {
	basename := filepath.Base(path)

	extraParams := map[string]string{
		"action":   "upload",
		"filename": basename,
		"token":    csrftoken,
		"format":   "json",
	}

	body, err := postFileRequest(cli, extraParams, "file", path)
	if err != nil {
		log.Fatal(err)
	}

	return body, nil
}

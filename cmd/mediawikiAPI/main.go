package main

import (
	"log"
	"net/http"
	"net/http/cookiejar"

	"github.com/k1rnt/mediawikiAPI/handler"
	"github.com/k1rnt/mediawikiAPI/pkg/api"
)

func setClient() *http.Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	// timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Jar: jar,
		// Timeout: timeout,
	}
	return client
}

func main() {
	cli := setClient()
	logintoken := api.GetLoginToken(cli).Query.Tokens.Logintoken
	api.GetLoginRequest(cli, logintoken)
	csrftoken := api.CsrfToken(cli).Query.Tokens.Csrftoken

	// wiki一括ページ作成
	if wiki_path != "" {
		if err := handler.UploadPageHandler(cli, csrftoken, wiki_path); err != nil {
			log.Fatal(err)
		}
	}

	// 画像一括投稿
	if images_path != "" {
		if err := handler.UploadImageHandler(cli, csrftoken, images_path); err != nil {
			log.Fatal(err)
		}
	}
}

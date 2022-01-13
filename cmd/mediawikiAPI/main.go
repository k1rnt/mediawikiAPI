package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"

	"github.com/k1rnt/mediawikiAPI/pkg/api"
	"github.com/k1rnt/mediawikiAPI/pkg/utils"
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

	// example#1
	// files, err := utils.FindWikiFiles(wiki_path)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, path := range files {
	// 	if _, err := api.EditRequest(cli, csrftoken, path); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("success %s\n", path)
	// }

	// example#2
	images, err := utils.FindImageFiles(images_path)
	if err != nil {
		log.Fatal(err)
	}
	for _, path := range images {
		body, err := api.Upload(cli, csrftoken, path)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", body)
		fmt.Printf("success %s\n", path)
	}
}

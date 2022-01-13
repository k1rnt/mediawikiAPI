package main

import (
	"flag"
	"log"
)

var (
	wiki_path   string
	images_path string
)

func init() {
	setFlag()
	// handle()
}

func setFlag() {
	flag.StringVar(&wiki_path, "wiki", "", "Select the directory containing the mediawiki files.")
	flag.StringVar(&images_path, "images", "", "Select the directory containing the imege files.")
	flag.Parse()
}

func handle() {
	readme := [...]string{wiki_path, images_path}
	for _, flagVar := range readme {
		if flagVar == "" {
			flag.PrintDefaults()
			log.Fatal("readme :-)")
		}
	}
}

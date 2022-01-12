package main

import (
	"flag"
	"log"
)

var (
	dir string
)

func init() {
	setFlag()
	handle()
}

func setFlag() {
	flag.StringVar(&dir, "dir", "", "Select the directory containing the mediawiki files.")
	flag.Parse()
}

func handle() {
	if dir == "" {
		flag.PrintDefaults()
		log.Fatal("readme :-)")
	}
}

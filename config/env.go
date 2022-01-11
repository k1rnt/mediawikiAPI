package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Endpoint = ""
	User     = ""
	Pass     = ""
)

func Env_load() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Endpoint = os.Getenv("API_ENDPOINT")
	User = os.Getenv("BOT_USER")
	Pass = os.Getenv("BOT_PASS")
}

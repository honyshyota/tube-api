package main

import (
	"flag"
	"log"

	apiserver "github.com/honyshyota/tube-api-go/internal/app/apiserver"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

}

func main() {
	flag.Parse()

	apiserver.Start()
}

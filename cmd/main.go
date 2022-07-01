package main

import (
	"flag"
	"log"
	"os"

	apiserver "github.com/honyshyota/tube-api-go/internal/app/apiserver"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

}

var (
	maxResults = flag.String("max-results", "10", "Max YouTube results")
)

func main() {
	flag.Parse()

	os.Setenv("MAX_RESULTS", *maxResults)

	apiserver.Start()
}

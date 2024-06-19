package config

import (
	"encoding/base64"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	ServerPort   = os.Getenv("SERVER_PORT")
	LokiEndpoint = os.Getenv("LOKI")
	BasicAuth    = "Basic " + base64.StdEncoding.EncodeToString([]byte(os.Getenv("BASIC_AUTH_USERNAME")+":"+os.Getenv("BASIC_AUTH_PASSWORD")))
)

// print the environment variables
func PrintEnv() {
	// Exit 1 when the environment variables are not set
	if LokiEndpoint == "" || BasicAuth == "" {
		println("LOKI, BASIC_AUTH_USERNAME, BASIC_AUTH_PASSWORD must be set")
		os.Exit(1)
	}

	print := func() {
		println("SERVER_PORT: ", ServerPort)
		println("LOKI: ", LokiEndpoint)
		println("BASIC_AUTH: ", BasicAuth)
	}
	print()
}

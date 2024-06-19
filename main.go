package main

import (
	"flag"
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	"gitlab.com/t9963/log-proxy-server/pkg/server"
)

var (
	version = "1.0.0" // initial version
)

func main() {
	versionFlag := flag.Bool("version", false, "Show version number")
	flag.Parse()

	if *versionFlag {
		fmt.Println("Version:", version)
		return
	}

	server.Start()
}

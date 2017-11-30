package main

import (
	"os"

	"github.com/LIANGTJ/cloudgo-data/service"
	flag "github.com/spf13/pflag"
)

const (
	PORT string = "8081"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	server := service.NewServer()
	server.Run(":" + port)
}

package main

import (
	"log"
	"os"

	"github.com/Tan4ek/openvpn-http-api/config"
	"github.com/Tan4ek/openvpn-http-api/server"
)

func initCAPassFile(CAPrivateKeyPass string) {
	f, err := os.Create("capassfile")
	if err != nil {
		log.Fatal(err)
	}

	f.WriteString(CAPrivateKeyPass)
	f.Close()
}

func main() {
	сonf := config.LoadConfig()

	initCAPassFile(сonf.CAPrivateKeyPass)

	server.Run(сonf.Server.Port)
}

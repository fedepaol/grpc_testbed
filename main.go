package main

import (
	"flag"

	"github.com/fedepaol/grpcsample/client"
	"github.com/fedepaol/grpcsample/server"
)

func main() {
	serverHost := flag.String("server", "127.0.0.1:2525", "the server host / port")
	serverPort := flag.Int("port", 2525, "the port the server start listening from")
	clientMode := flag.Bool("client", false, "client mode")

	if *clientMode {
		client.Run(*serverHost)
	} else {
		server.Run(*serverPort)
	}
}

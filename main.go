package main

import (
	"flag"
	"log"

	"github.com/logeshsuresh/rediscovery/config"
	"github.com/logeshsuresh/rediscovery/server"
)

func init() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the rediscovery server")
	flag.IntVar(&config.Port, "port", 8379, "port for the rediscovery server")
	flag.Parse()
}

func main() {
	log.Println("Rediscovery underway: Unlocking the hidden powers of Redis from scratch! 🔐🔥")
	server.RunSyncTCPServer()
}
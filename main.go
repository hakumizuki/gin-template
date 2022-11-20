package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hakumizuki/gin-template/config"
	"github.com/hakumizuki/gin-template/db"
	"github.com/hakumizuki/gin-template/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Start()
}

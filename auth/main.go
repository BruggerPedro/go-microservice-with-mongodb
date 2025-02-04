package main

import (
	"flag"
	"log"
	"github.com/joho/godotenv"
	"microservice-go/db"
)

var (
	local bool
)

func init() {
	flag.BoolVar(&local, "local", true, "Run server in local mode")
	flag.Parse()
}

func main() {
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Panicln(err)
		}
	}

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Panicln(err)
	}
	
	defer conn.Close()


}
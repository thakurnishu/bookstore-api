package main

import (
	"flag"
	"log"

	"github.com/thakurnishu/bookstore-api/api"
	"github.com/thakurnishu/bookstore-api/storage"
)

func main() {

	listenAdrr := flag.String("listenAdrr", "3000", "Port where server will listen")
	flag.Parse()

	// store, err := storage.NewPostgresStore()
	store, err := storage.NewMySQLStore()
	if err != nil {
		log.Fatalln(err)
	}
	if err = store.Init(); err != nil {
		log.Fatalln(err)
	}

	server := api.NewServer(*listenAdrr, store)
	server.Run()

}

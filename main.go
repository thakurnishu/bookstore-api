package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thakurnishu/mysql-api/pkg/routes"
)

const (
	defaultPort = "8000"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	log.Printf("Listen on Post : %s", defaultPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", defaultPort), r)
	if err != nil {
		log.Printf("ERROR: listening to port\n%s\n\n", err.Error())
	}
}

package routes

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/thakurnishu/bookstore-api/pkg/models"
	"github.com/thakurnishu/bookstore-api/pkg/utils"
)

const (
	// Default API-Path
	defaultAPIPath = "/apis/v1/books"
)

var RegisterRoutes = func(r *mux.Router) {

	log.Println("Register Routes")

	r.HandleFunc(getAPIPath(), models.GetBooks).Methods("GET")
}

func getAPIPath() string {
	apiPath := utils.GetEnvOrDefault("API_PATH", defaultAPIPath)

	utils.GetEnvOrDefault("API_PATH", defaultAPIPath)

	return apiPath
}

package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/thakurnishu/bookstore-api/storage"
	"github.com/thakurnishu/bookstore-api/utils"
)

type Server struct {
	ListenAddr string
	store      storage.Storage
}

type ServerError struct {
	Error string `json:"error"`
}

func NewServer(listenAdrr string, store storage.Storage) *Server {
	return &Server{
		ListenAddr: fmt.Sprintf(":%s", listenAdrr),
		store:      store,
	}
}

func (s *Server) Run() {

	router := httprouter.New()

	router.HandlerFunc("GET", "/book", utils.CustomHTTPHandleFunc(s.HandleGetBook))
	router.HandlerFunc("GET", "/book/:title", utils.CustomHTTPHandleFunc(s.HandleGetBookByTitle))

	router.HandlerFunc("POST", "/book/register", utils.CustomHTTPHandleFunc(s.HandleRegisterBook))
	router.HandlerFunc("PUT", "/book/update/:title", utils.CustomHTTPHandleFunc(s.HandleUpdateBookByTitle))
	router.HandlerFunc("DELETE", "/book/remove/:title", utils.CustomHTTPHandleFunc(s.HandleDeleteBookByTitle))

	log.Printf("Server Listen on Port %s\n", s.ListenAddr)
	log.Fatalln(http.ListenAndServe(s.ListenAddr, router))
}

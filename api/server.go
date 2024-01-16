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

func NewServer(listenAdrr string, store storage.Storage) *Server {
	return &Server{
		ListenAddr: fmt.Sprintf(":%s", listenAdrr),
		store:      store,
	}
}

func (s *Server) Run() {

	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/book/register", utils.CustomHTTPHandleFunc(s.HandleRegisterBook))
	router.HandlerFunc(http.MethodGet, "/book", utils.CustomHTTPHandleFunc(s.HandleGetBook))

	log.Printf("Server Listen on Port %s\n", s.ListenAddr)
	log.Fatalln(http.ListenAndServe(s.ListenAddr, router))
}

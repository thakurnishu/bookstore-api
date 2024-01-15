package api

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) HandleRegisterBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	log.Println("Registered")
	return nil
}

func (s *Server) HandleGetBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	log.Println("Got Books")
	return nil
}

func (s *Server) HandleGetBookByTitle(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	log.Println("Got Book")
	return nil
}

func (s *Server) HandleUpdateBookByTitle(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	log.Println("Updated")
	return nil
}

func (s *Server) HandleDeleteBookByTitle(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	log.Println("Deleted")
	return nil
}

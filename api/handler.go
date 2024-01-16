package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/thakurnishu/bookstore-api/types"
	"github.com/thakurnishu/bookstore-api/utils"
)

func (s *Server) HandleRegisterBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {

	defer r.Body.Close()

	bookReq := new(types.BookRequest)
	if err := json.NewDecoder(r.Body).Decode(bookReq); err != nil {
		return err
	}
	bookReq = utils.ConvertAllToTitle(bookReq)

	book := utils.NewBook(bookReq)
	if err := s.store.AddBook(book); err != nil {
		return err
	}

	bookResp := utils.BookResp(book)

	return utils.WriteJSON(w, http.StatusCreated, bookResp)
}

func (s *Server) HandleGetBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {

	if hasQueryParameter := r.URL.Query().Has("title"); !hasQueryParameter {
		books, err := s.store.GetBook()
		if err != nil {
			return err
		}

		return utils.WriteJSON(w, http.StatusAccepted, books)
	} else {
		return s.HandleGetBookWithTitleQuery(w, r, params)
	}

}

func (s *Server) HandleGetBookWithTitleQuery(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {

	title := utils.GetQueryValTitle(r, "title")

	bookResp, err := s.store.GetBookByTitle(title)
	if err != nil {
		return err
	}

	if bookResp.Title == "" {
		return fmt.Errorf("book doesn't exist")
	}

	return utils.WriteJSON(w, http.StatusAccepted, bookResp)
}

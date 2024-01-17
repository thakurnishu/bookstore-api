package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/thakurnishu/bookstore-api/types"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func CustomHTTPHandleFunc(f types.CustomHandleFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		params := httprouter.ParamsFromContext(r.Context())

		if err := f(w, r, params); err != nil {
			log.Println(err)
			WriteJSON(w, http.StatusBadRequest, types.ApiError{
				Error: err.Error(),
			})
		}

	}
}

func ConvertAllToTitle(book *types.BookRequest) *types.BookRequest {

	book.Author = title(book.Author)
	book.Title = title(book.Title)
	book.Publication = title(book.Publication)

	return book
}

func title(str string) string {

	casesr := cases.Title(language.English)
	str = strings.ToLower(str)

	return casesr.String(str)
}

func GetQueryValTitle(r *http.Request, key string) string {

	casesr := cases.Title(language.English)
	queryVal := r.URL.Query().Get(key)
	queryVal = casesr.String(strings.ReplaceAll(queryVal, "-", " "))

	return queryVal
}

package utils

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type customHandleFunc func(http.ResponseWriter, *http.Request, httprouter.Params) error

type customApiError struct {
	Error string `json:"error"`
}

func CustomHTTPHandleFunc(f customHandleFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		params := httprouter.ParamsFromContext(r.Context())

		if err := f(w, r, params); err != nil {
			WriteJSON(w, http.StatusBadRequest, customApiError{
				Error: err.Error(),
			})
		}

	}
}

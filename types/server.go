package types

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CustomHandleFunc func(http.ResponseWriter, *http.Request, httprouter.Params) error

type ApiError struct {
	Error string `json:"error"`
}

package storage

import (
	"github.com/thakurnishu/bookstore-api/types"
)

type Storage interface {
	AddBook(*types.Book) error
	GetBook() ([]*types.BookResponse, error)
	GetBookByTitle(string) (*types.BookResponse, error)
}

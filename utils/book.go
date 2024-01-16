package utils

import (
	"database/sql"
	"time"

	"github.com/thakurnishu/bookstore-api/types"
)

func NewBook(book *types.BookRequest) *types.Book {
	return &types.Book{
		Title:       book.Title,
		Author:      book.Author,
		Publication: book.Publication,
		Isbn:        book.Isbn,
		Available:   book.Available,
		Added_At:    time.Now(),
	}
}

func BookResp(book *types.Book) *types.BookResponse {
	return &types.BookResponse{
		Title:       book.Title,
		Author:      book.Author,
		Publication: book.Publication,
		Isbn:        book.Isbn,
		Available:   book.Available,
		Added_At:    book.Added_At,
	}
}

func ScanBook(rows *sql.Rows) (*types.Book, error) {
	book := new(types.Book)
	if err := rows.Scan(
		&book.Id,
		&book.Available,
		&book.Added_At,
		&book.Title,
		&book.Author,
		&book.Publication,
		&book.Isbn,
	); err != nil {
		return nil, err
	}

	return book, nil
}

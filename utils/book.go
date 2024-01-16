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

func ScanBookPostgres(rows *sql.Rows) (*types.Book, error) {
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

func ScanBookMySQL(rows *sql.Rows) (*types.Book, error) {

	var added_at []uint8

	book := new(types.Book)
	if err := rows.Scan(
		&book.Id,
		&book.Available,
		&added_at,
		&book.Title,
		&book.Author,
		&book.Publication,
		&book.Isbn,
	); err != nil {
		return nil, err
	}

	parsedTime, err := time.Parse("2006-01-02 15:04:05", string(added_at))
	if err != nil {
		return nil, err
	}

	book.Added_At = parsedTime

	return book, nil
}

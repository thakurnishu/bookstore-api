package storage

import (
	"fmt"
	"time"

	"github.com/thakurnishu/bookstore-api/types"
	"github.com/thakurnishu/bookstore-api/utils"
)

type Storage interface {
	AddBook(*types.Book) error
	GetBook() ([]*types.BookResponse, error)
	GetBookByTitle(string) (*types.BookResponse, error)
}

func (db *PostgresStore) AddBook(book *types.Book) error {
	query := `
	INSERT INTO book (
			title,
			author, 
			publication, 
			isbn, 
			available,
			added_at
		)
        values ($1, $2, $3, $4, $5, $6)
	`

	_, err := db.Exec(query, book.Title, book.Author, book.Publication, book.Isbn, book.Available, time.Now())
	if err != nil {
		// return fmt.Errorf("failed to create account")
		return err
	}
	return nil
}

func (db *PostgresStore) GetBook() ([]*types.BookResponse, error) {

	rows, err := db.Query("SELECT * FROM book")
	if err != nil {
		return nil, err
	}

	books := []*types.BookResponse{}

	for rows.Next() {
		book, err := utils.ScanBook(rows)
		if err != nil {
			return nil, err
		}
		bookResp := utils.BookResp(book)

		books = append(books, bookResp)
	}

	return books, nil
}

func (db *PostgresStore) GetBookByTitle(title string) (*types.BookResponse, error) {

	bookResp := types.BookResponse{}

	rows, err := db.Query("SELECT * FROM book where title = $1", title)
	if err != nil {
		return nil, fmt.Errorf("book doesn't exist")
	}

	for rows.Next() {
		book, err := utils.ScanBook(rows)
		if err != nil {
			return nil, fmt.Errorf("book doesn't exist")
		}

		bookResp = *utils.BookResp(book)
		break
	}
	return &bookResp, err
}

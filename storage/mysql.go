package storage

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
	"github.com/thakurnishu/bookstore-api/types"
	"github.com/thakurnishu/bookstore-api/utils"
)

type MySQLStore struct {
	*sql.DB
}

func NewMySQLStore() (*MySQLStore, error) {
	godotenv.Load()

	var (
		dbname     = os.Getenv("DB_NAME")
		dbuser     = os.Getenv("DB_USER")
		dbpassword = os.Getenv("DB_PASSWORD")
		dbhost     = os.Getenv("DB_HOST")
		dbport     = os.Getenv("DB_PORT")
		uri        = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpassword, dbhost, dbport, dbname)
	)

	db, err := sql.Open("mysql", uri)
	if err != nil {
		return nil, fmt.Errorf("opening sql connection: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping sql db: %s ", err.Error())
	}

	return &MySQLStore{
		DB: db,
	}, nil
}

func (db *MySQLStore) Init() error {
	return db.createAccountTable()
}

func (db *MySQLStore) createAccountTable() error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS book (
  		id INT AUTO_INCREMENT PRIMARY KEY,
  		available INT,
  		added_at TIMESTAMP,
  		title VARCHAR(255) NOT NULL,
  		author VARCHAR(255) NOT NULL,
  		publication VARCHAR(255) NOT NULL,
  		isbn BIGINT UNIQUE NOT NULL
	);`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create table: %s", err.Error())
	}
	return nil
}

func (db *MySQLStore) AddBook(book *types.Book) error {
	query := `
	INSERT INTO book (
			title,
			author, 
			publication, 
			isbn, 
			available,
			added_at
		)
        values (?, ?, ?, ?, ?, ?)
	`

	_, err := db.Exec(query, book.Title, book.Author, book.Publication, book.Isbn, book.Available, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (db *MySQLStore) GetBook() ([]*types.BookResponse, error) {

	rows, err := db.Query("SELECT * FROM book")
	if err != nil {
		return nil, err
	}

	books := []*types.BookResponse{}

	for rows.Next() {
		book, err := utils.ScanBookMySQL(rows)
		if err != nil {
			return nil, err
		}

		bookResp := utils.BookResp(book)

		books = append(books, bookResp)
	}

	return books, nil
}

func (db *MySQLStore) GetBookByTitle(title string) (*types.BookResponse, error) {

	bookResp := types.BookResponse{}

	rows, err := db.Query("SELECT * FROM book where title = ?", title)
	if err != nil {
		return nil, fmt.Errorf("book doesn't exist")
	}

	for rows.Next() {
		book, err := utils.ScanBookMySQL(rows)
		if err != nil {
			return nil, fmt.Errorf("book doesn't exist")
		}

		bookResp = *utils.BookResp(book)
		break
	}
	return &bookResp, err
}

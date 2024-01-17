package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/thakurnishu/bookstore-api/types"
	"github.com/thakurnishu/bookstore-api/utils"
)

type PostgresStore struct {
	*sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	godotenv.Load()

	var (
		dbname     = os.Getenv("DB_NAME")
		dbuser     = os.Getenv("DB_USER")
		dbpassword = os.Getenv("DB_PASSWORD")
		dbhost     = os.Getenv("DB_HOST")
		dbport     = os.Getenv("DB_PORT")
	)
	if dbport == "" {
		dbport = "5432"
	}
	uri := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", dbuser, dbname, dbpassword, dbhost, dbport)

	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("opening postgres connection: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to ping postgres db: %s ", err.Error())
	}

	return &PostgresStore{
		DB: db,
	}, nil
}

func (db *PostgresStore) Init() error {
	return db.createAccountTable()
}

func (db *PostgresStore) createAccountTable() error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS "book" (
  		"id" serial PRIMARY KEY,
  		"available" int,
  		"added_at" timestamp,
  		"title" varchar NOT NULL,
  		"author" varchar NOT NULL,
  		"publication" varchar NOT NULL,
  		"isbn" bigint UNIQUE NOT NULL
	);`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to create table: %s", err.Error())
	}
	return nil
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
		log.Println(err)
		// return fmt.Errorf("failed to create account")
		return err
	}
	return nil
}

func (db *PostgresStore) GetBook() ([]*types.BookResponse, error) {

	rows, err := db.Query("SELECT * FROM book")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	books := []*types.BookResponse{}

	for rows.Next() {
		book, err := utils.ScanBookPostgres(rows)
		if err != nil {
			log.Println(err)
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
		log.Println(err)
		return nil, fmt.Errorf("book doesn't exist")
	}

	for rows.Next() {
		book, err := utils.ScanBookPostgres(rows)
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("book doesn't exist")
		}

		bookResp = *utils.BookResp(book)
		break
	}
	return &bookResp, err
}

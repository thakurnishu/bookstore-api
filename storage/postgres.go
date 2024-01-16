package storage

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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
		uri        = fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", dbuser, dbname, dbpassword, dbhost, dbport)
	)

	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, fmt.Errorf("opening postgres connection: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
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
		return fmt.Errorf("failed to create table: %s", err.Error())
	}
	return nil
}

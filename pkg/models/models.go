package models

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/thakurnishu/mysql-api/pkg/db"
)

type Book struct {
	Id   string
	Name string
	Isbn string
}

var databaseConnection *sql.DB

func getdatabaseConnection() {
	databaseConnection = db.GetDB()
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("From Get Books")

	getdatabaseConnection()
	rows, err := databaseConnection.Query("select * from books")
	if err != nil {
		log.Printf("ERROR: getting books from database\n%s\n\n", err.Error())
	}

	for rows.Next() {

	}
}

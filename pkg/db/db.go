package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/thakurnishu/bookstore-api/pkg/utils"
)

var (
	connectionDB *sql.DB
)

const (
	// Default Database Variable
	defaultHost     = "localhost:3306"
	defaultPass     = "nishant"
	defaultName     = "library"
	defaultUser     = "root"
	defaultProtocol = "tcp"
)

type dataBase struct {
	// Host contain host-ip:port
	Host     string
	Name     string
	Pass     string
	User     string
	Protocol string
}

func init() {

	log.Println("init")
	dbInfo := dataBase{}
	Info := dbInfo.GetDBConfig()
	openConnection(Info)

}

func (dbString *dataBase) GetDBConfig() *dataBase {

	db := dataBase{
		Host:     utils.GetEnvOrDefault("DB_HOST", defaultHost),
		Pass:     utils.GetEnvOrDefault("DB_PASS", defaultPass),
		Name:     utils.GetEnvOrDefault("DB_NAME", defaultName),
		User:     utils.GetEnvOrDefault("DB_USER", defaultUser),
		Protocol: utils.GetEnvOrDefault("DB_PROTOCOL", defaultProtocol),
	}
	return &db
}

func openConnection(dbInfo *dataBase) {

	dns := fmt.Sprintf("%s:%s@%s(%s)/%s", (*dbInfo).User, (*dbInfo).Pass, (*dbInfo).Protocol, (*dbInfo).Host, (*dbInfo).Name)

	dbConnection, err := sql.Open("mysql", dns)
	if err != nil {
		log.Printf("ERROR: opening connection to database\n%s\n\n", err.Error())
	}

	err = dbConnection.Ping()
	if err != nil {
		log.Printf("ERROR: connecting to database failed\n%s\n\n", err.Error())
	}

	connectionDB = dbConnection

}

func GetDB() *sql.DB {
	return connectionDB
}

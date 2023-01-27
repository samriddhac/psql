package pghelper

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	HostName string = ""
	Port     int    = 5432
	UserName string = ""
	Password string = ""
	Database string = ""
)

type User struct {
	ID string
	Username string
}

func openConnection() (*sqlx.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HostName, Port, UserName, Password, Database)
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ExecuteQueryUsersAll(query string, args ...interface{}) (*[]User, error) {
	db, err := openConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	response := []User{}
	dataerr := db.Select(&response, query, args...)
	if dataerr != nil {
		return nil, dataerr
	}
	return &response, nil
}
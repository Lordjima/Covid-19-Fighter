package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" //jutified
	"log"
)

var db *sql.DB

const (
	host     = "0.0.0."
	port     = 5432
	user     = "unicorn_user"
	password = "magical_password"
	dbname   = "postgres"
)

//DatabaseInit init db
func DatabaseInit() {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	// Create Table cars if not exists
	createCarsTable()
}

func createCarsTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS patients(id serial,patientID varchar(255), patientAddr varchar(255), decryptKey varchar(255), created_at timestamp default NULL, updated_at timestamp default NULL, constraint pk primary key(id))")

	if err != nil {
		log.Fatal(err)
	}
}

//Db Getter for db var
func Db() *sql.DB {
	return db
}

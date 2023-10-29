package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func DBInitialization() {
	var (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "pritam"
		dbname   = "postgres"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	cleint, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Tidak Konek DB Errornya : %s", err)
	}
	Db = cleint
	log.Println("Database connection successful")

}

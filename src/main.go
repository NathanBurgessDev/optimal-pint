package main

import (
	"log"
	"net/http"
	"optimal-pint/src/internal/service"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sqlx.Connect("sqlite3", "../Spooninit.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	service := service.NewService(db)

	http.HandleFunc("/getPubs", service.AllPubs)

	http.ListenAndServe(":8080", nil)
}

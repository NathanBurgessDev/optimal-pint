package main

import (
	"net/http"
	"optimal-pint/src/internal/fetcher"
	"optimal-pint/src/internal/service"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := sqlx.MustConnect("sqlite3", "test.db")

	db.MustExec(`CREATE TABLE "Pubs" (
		"ID"		INTEGER NOT NULL,
		"PubID" 	INTEGER NOT NULL,
		"PubName"	TEXT NOT NULL,
		"Longitude"	REAL NOT NULL,
		"Latitude"	REAL NOT NULL,
		"City"		TEXT,
		PRIMARY KEY("ID" AUTOINCREMENT)
	);`)

	f := fetcher.New(db)

	pubService := service.NewService(db)

	http.HandleFunc("/getPubs", pubService.AllPubs)

	// http.ListenAndServe(":8080", nil)
}

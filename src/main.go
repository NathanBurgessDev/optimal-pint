package main

import (
	"fmt"
	"net/http"
	"optimal-pint/src/internal/service"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := sqlx.MustConnect("sqlite3", "../test.db")

	// db.MustExec(`CREATE TABLE "Pubs" (
	// 	"ID"		INTEGER NOT NULL,
	// 	"PubID" 	INTEGER NOT NULL,
	// 	"PubName"	TEXT NOT NULL,
	// 	"Longitude"	REAL NOT NULL,
	// 	"Latitude"	REAL NOT NULL,
	// 	"City"		TEXT,
	// 	PRIMARY KEY("ID" AUTOINCREMENT)
	// );`)

	// f := fetcher.New(db)

	pubService := service.NewService(db)

	http.HandleFunc("/tfarndai", pubService.AllPubs)

	fmt.Println("Server starting on :8090")
	http.ListenAndServe(":8090", nil)
}

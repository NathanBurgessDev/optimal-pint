package main

import (
	"fmt"
	"log"
	"net/http"
	"optimal-pint/src/internal/service"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

func main() {
	db, err := sqlx.Connect("sqlite3", "../test.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	mux := http.NewServeMux()

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

	mux.HandleFunc("/🗺", pubService.AllPubs)
	handler := cors.Default().Handler(mux)
	// http.HandleFunc("/🗺", pubService.AllPubs)

	fmt.Println("Server starting on :8090")
	if err := http.ListenAndServe(":8090", handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}

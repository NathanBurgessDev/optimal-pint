package service

import (
	"fmt"
	"log"
	"net/http"
	"optimal-pint/src/internal/pubInfo"

	"github.com/jmoiron/sqlx"
)

type Service struct {
	db *pubInfo.DB
}

func NewService(db *sqlx.DB) *Service {
	return &Service{pubInfo.NewDB(db)}
}

func (s *Service) AllPubs(w http.ResponseWriter, r *http.Request) {
	data, err := s.db.GetAllPubs()
	if err != nil {
		err = fmt.Errorf("error retrieving pubs: %v", err)
		log.Fatalf("%v", err.Error())
		http.Error(w, "Failed to retrieve pubs", http.StatusInternalServerError)
		return
	}
	// fmt.Print(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

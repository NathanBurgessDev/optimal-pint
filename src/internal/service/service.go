package service

import (
	"fmt"
	"log"
	"net/http"
	"optimal-pint/src/internal/pubInfo"

	"github.com/jmoiron/sqlx"
)

const internalServerStr = "Wuh Woh OwO *gasps* theres been a fucky wucky!"

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
		http.Error(w, internalServerStr, http.StatusInternalServerError)
		return
	}
	// fmt.Print(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (s *Service) AllDrinks(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	data, err := s.db.GetAllDrinks(idString)
	if err != nil {
		err = fmt.Errorf("error retrieving drinks: %v", err)
		log.Fatalf("%v", err.Error())
		http.Error(w, internalServerStr, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

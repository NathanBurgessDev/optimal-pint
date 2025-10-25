package pubInfo

import (
	"encoding/json"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	db *sqlx.DB
}

type Pub struct {
	ID        int     `db:"PubID"`
	Name      string  `db:"PubName"`
	Owner     string  `db:"PubOwner"`
	Longitude float64 `db:"Longitude"`
	Latitude  float64 `db:"Latitude"`
	City      string  `db:"City"`
}

type PubJson struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Owner     string  `json:"owner"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	City      string  `json:"city"`
}

func NewDB(db *sqlx.DB) *DB {
	return &DB{db: db}
}

func (s *DB) GetAllPubs() ([]byte, error) {
	var pubs []Pub
	err := s.db.Select(&pubs, "SELECT * FROM Pubs")
	if err != nil {
		return nil, err
	}

	var pubJsonList []PubJson
	for _, pub := range pubs {
		pubJson := PubJson(pub)
		pubJsonList = append(pubJsonList, pubJson)
	}

	jsonData, err := json.Marshal(pubJsonList)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

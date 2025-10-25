package pubInfo

import (
	"encoding/json"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	db *sqlx.DB
}

type Pub struct {
	ID        int     `db:"ID"`
	SpoonsID  int     `db:"PubID"`
	Name      string  `db:"PubName"`
	Longitude float64 `db:"Longitude"`
	Latitude  float64 `db:"Latitude"`
	City      string  `db:"City"`
}

type PubJson struct {
	ID        int     `json:"id"`
	SpoonsID  int     `json:"spoonsId"`
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	City      string  `json:"city"`
}

type Drink struct {
	ID        int     `db:"ID"`
	PubID     int     `db:"PubID"`
	DrinkName string  `db:"DrinkName"`
	Units     float64 `db:"Units"`
	Price     float64 `db:"Price"`
	Amount    int     `db:"Amount"`
	Category  string  `db:"Category"`
}

type DrinkJson struct {
	ID        int     `json:"id"`
	PubID     int     `json:"pubId"`
	DrinkName string  `json:"drinkName"`
	Units     float64 `json:"units"`
	Price     float64 `json:"price"`
	Amount    int     `json:"amount"`
	Category  string  `json:"category"`
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

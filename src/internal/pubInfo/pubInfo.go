package pubInfo

import (
	"encoding/json"
	"fmt"

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
	ID         int     `db:"ID"`
	PubID      int     `db:"PubID"`
	DrinkName  string  `db:"DrinkName"`
	Units      float64 `db:"Units"`
	Price      float64 `db:"Price"`
	Amount     int     `db:"Amount"`
	Category   string  `db:"Category"`
	Optimality float64 `db:"Optimality"`
	HasDeal    bool    `db:"HasDeal"`
}

type DrinkJson struct {
	ID         int     `json:"id"`
	PubID      int     `json:"pubId"`
	DrinkName  string  `json:"drinkName"`
	Units      float64 `json:"units"`
	Price      float64 `json:"price"`
	Amount     int     `json:"amount"`
	Category   string  `json:"category"`
	Optimality float64 `json:"poundsPerUnit"`
	HasDeal    bool    `json:"hasDeal"`
}

func NewDB(db *sqlx.DB) *DB {
	return &DB{db: db}
}

func (s *DB) GetAllPubs() ([]byte, error) {
	var pubs []Pub
	err := s.db.Select(&pubs, "SELECT * FROM Pubs")
	if err != nil {
		err = fmt.Errorf("error selecting pubs: %v", err)
		return nil, err
	}

	var pubJsonList []PubJson
	for _, pub := range pubs {
		pubJson := PubJson(pub)
		pubJsonList = append(pubJsonList, pubJson)
	}

	jsonData, err := json.Marshal(pubJsonList)
	if err != nil {
		err = fmt.Errorf("error marshalling pubs to JSON: %v", err)
		return nil, err
	}
	return jsonData, nil
}

func (s *DB) GetAllDrinks(pubID string) ([]byte, error) {
	var drinks []Drink
	err := s.db.Select(&drinks, "SELECT * FROM Drinks WHERE PubID = ? and Optimality < 999 ORDER BY Optimality ASC", pubID)
	if err != nil {
		err = fmt.Errorf("error selecting drinks: %v", err)
		return nil, err
	}

	var drinkJsonList []DrinkJson
	for _, drink := range drinks {
		drinkJson := DrinkJson(drink)
		drinkJsonList = append(drinkJsonList, drinkJson)
	}

	jsonData, err := json.Marshal(drinkJsonList)
	if err != nil {
		err = fmt.Errorf("error marshalling drinks to JSON: %v", err)
		return nil, err
	}
	return jsonData, nil
}

func (s *DB) GetAllDrinksWithDeals(pubID string) ([]byte, error) {
	var drinks []Drink
	err := s.db.Select(&drinks, "SELECT * FROM Drinks WHERE PubID = ? AND HasDeal = 1 and Optimality < 999 ORDER BY Optimality ASC", pubID)
	if err != nil {
		err = fmt.Errorf("error selecting drinks with deals: %v", err)
		return nil, err
	}

	var drinkJsonList []DrinkJson
	for _, drink := range drinks {
		drinkJson := DrinkJson(drink)
		drinkJsonList = append(drinkJsonList, drinkJson)
	}

	jsonData, err := json.Marshal(drinkJsonList)
	if err != nil {
		err = fmt.Errorf("error marshalling drinks with deals to JSON: %v", err)
		return nil, err
	}
	return jsonData, nil
}

func (s *DB) GetTopDrinks() ([]byte, error) {
	var drinks []Drink
	err := s.db.Select(&drinks, "SELECT ID, PubID, DrinkName, Units, Price, Amount, Category, Optimality, HasDeal FROM (SELECT *, ROW_NUMBER() OVER (PARTITION BY DrinkName ORDER BY Optimality ASC) AS rn FROM Drinks WHERE Optimality < 999) ranked WHERE rn = 1 ORDER BY Optimality ASC LIMIT 100;")
	if err != nil {
		err = fmt.Errorf("error selecting top drinks: %v", err)
		return nil, err
	}

	var drinkJsonList []DrinkJson
	for _, drink := range drinks {
		drinkJson := DrinkJson(drink)
		drinkJsonList = append(drinkJsonList, drinkJson)
	}

	jsonData, err := json.Marshal(drinkJsonList)
	if err != nil {
		err = fmt.Errorf("error marshalling top drinks to JSON: %v", err)
		return nil, err
	}
	return jsonData, nil
}

func (s *DB) GetPubByID(pubID string) ([]byte, error) {
	var pub Pub
	err := s.db.Get(&pub, "SELECT * FROM Pubs WHERE PubID = ?", pubID)
	if err != nil {
		err = fmt.Errorf("error selecting pub by ID: %v", err)
		return nil, err
	}
	pubJson := PubJson(pub)
	data, err := json.Marshal(pubJson)
	if err != nil {
		err = fmt.Errorf("error marshalling pub to JSON: %v", err)
		return nil, err
	}
	return data, nil
}

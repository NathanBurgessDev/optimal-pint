package fetcher

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Fetcher struct {
	client *http.Client
	db     *sqlx.DB
}

func New(db *sqlx.DB) *Fetcher {
	return &Fetcher{
		client: &http.Client{},
		db:     db,
	}
}

func (f *Fetcher) Update() error {
	venues_res, err := f.client.Get("https://ca.jdw-apps.net/api/v0.1/venues")

	if err != nil {
		return err
	}

	var venues []Venue

	if err := json.NewDecoder(venues_res.Body).Decode(&venues); err != nil {
		return err
	}

}

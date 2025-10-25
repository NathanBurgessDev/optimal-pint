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
	req, err := http.NewRequest("GET", "https://ca.jdw-apps.net/api/v0.1/venues", nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer 1|SFS9MMnn5deflq0BMcUTSijwSMBB4mc7NSG2rOhqb2765466")

	venues_res, err := f.client.Do(req)

	if err != nil {
		return err
	}

	var venues VenuesResponse

	if err := json.NewDecoder(venues_res.Body).Decode(&venues); err != nil {
		return err
	}

	for _, venue := range venues.Data {
		_, err := f.db.Exec(
			"INSERT into Pubs (PubID, PubName, Longitude, Latitude, City) VALUES ($1, $2, $3, $4, $5)",
			venue.ID, venue.Name, venue.Address.Location.Longitude, venue.Address.Location.Latitude, venue.Address.Town,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

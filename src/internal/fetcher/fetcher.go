package fetcher

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

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

func (f *Fetcher) UpdateMenu(venueID int, salesAreaID int, menuID int) error {
	url := fmt.Sprintf("https://ca.jdw-apps.net/api/v0.1/jdw/venues/%d/sales-areas/%d/menus/%d", venueID, salesAreaID, menuID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer 1|SFS9MMnn5deflq0BMcUTSijwSMBB4mc7NSG2rOhqb2765466")

	res, err := f.client.Do(req)
	if err != nil {
		return err
	}

	var menu MenuDetailResponse

	if err := json.NewDecoder(res.Body).Decode(&menu); err != nil {
		return err
	}

	for _, category := range menu.Data.Categories {
		if category.Name == "Includes a drink" {
			continue
		}

		for _, itemGroup := range category.ItemGroups {
			for _, item := range itemGroup.Items {
				var units float64

				pattern := `([\d.]+)\s*units`
				re := regexp.MustCompile(pattern)

				matches := re.FindStringSubmatch(item.Description)
				if len(matches) >= 2 {
					units, _ = strconv.ParseFloat(matches[1], 64)
				}

				if len(item.Options.Portion.Options) == 0 {
					continue
				}

				maxOption := item.Options.Portion.Options[0]

				for _, option := range item.Options.Portion.Options {
					if option.Value.Price.Value > maxOption.Value.Price.Value {
						maxOption = option
					}
				}

				realPrice := maxOption.Value.Price.Value
				hasDeal := false

				if len(item.Options.Linked) != 0 {
					amount := 0
					dealPrice := 0.0

					link := item.Options.Linked[0]

					fmt.Sscanf(link.Name, "Any %d for £%f", &amount, &dealPrice)

					if amount != 0 && dealPrice != 0 {
						hasDeal = true
						realPrice = dealPrice / float64(amount)
					}
				}

				optimality := realPrice / units

				_, err := f.db.Exec(
					"INSERT into Drinks (DrinkName, PubId, Units, Price, Amount, Category, Optimality, HasDeal) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
					item.Name, venueID, units, realPrice, 67, category.Name, optimality, hasDeal,
				)

				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (f *Fetcher) UpdateBars(venueID int, salesArea SalesArea) error {
	url := fmt.Sprintf("https://ca.jdw-apps.net/api/v0.1/jdw/venues/%d/sales-areas/%d/menus?type=available", venueID, salesArea.ID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer 1|SFS9MMnn5deflq0BMcUTSijwSMBB4mc7NSG2rOhqb2765466")

	res, err := f.client.Do(req)
	if err != nil {
		return err
	}

	var menus MenusResponse

	if err := json.NewDecoder(res.Body).Decode(&menus); err != nil {
		return err
	}

	for _, menu := range menus.Data {
		if menu.Name == "Drinks" {
			return f.UpdateMenu(venueID, salesArea.ID, menu.ID)
		}
	}

	// return fmt.Errorf("Menu not found :(")
	return nil
}

func (f *Fetcher) UpdateVenue(venueID int) error {
	url := fmt.Sprintf("https://ca.jdw-apps.net/api/v0.1/venues/%d", venueID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer 1|SFS9MMnn5deflq0BMcUTSijwSMBB4mc7NSG2rOhqb2765466")

	res, err := f.client.Do(req)
	if err != nil {
		return err
	}

	var venueInfo VenueInfoResponse

	if err := json.NewDecoder(res.Body).Decode(&venueInfo); err != nil {
		return err
	}

	for _, salesArea := range venueInfo.Data.SalesAreas {
		if err := f.UpdateBars(venueID, salesArea); err != nil {
			return err
		}
	}

	return nil
}

func (f *Fetcher) Update() error {
	req, err := http.NewRequest("GET", "https://ca.jdw-apps.net/api/v0.1/venues", nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer 1|SFS9MMnn5deflq0BMcUTSijwSMBB4mc7NSG2rOhqb2765466") // Fuck you

	venues_res, err := f.client.Do(req)

	if err != nil {
		return err
	}

	var venues VenuesResponse

	if err := json.NewDecoder(venues_res.Body).Decode(&venues); err != nil {
		return err
	}

	for _, venue := range venues.Data {
		// if venue.Name != "The Joseph Else" && venue.Name != "The Gooseberry Bush" {
		// 	continue
		// }

		_, err := f.db.Exec(
			"INSERT into Pubs (PubID, PubName, Longitude, Latitude, City) VALUES ($1, $2, $3, $4, $5)",
			venue.VenueRef, venue.Name, venue.Address.Location.Longitude, venue.Address.Location.Latitude, venue.Address.Town,
		)

		if err != nil {
			return err
		}

		if err := f.UpdateVenue(venue.VenueRef); err != nil {
			return err
		}

		log.Printf("Added Venue %s", venue.Name)
	}

	return nil
}

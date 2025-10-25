package fetcher

type Address struct {
	Town     string   `json:"town"`
	Location Location `json:"location"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Venue struct {
	ID       int     `json:"id"`
	VenueRef int     `json:"venueRef"`
	Name     string  `json:"name"`
	Address  Address `json:"address"`
}

type VenuesResponse struct {
	Data []Venue `json:"data"`
}

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

type SalesArea struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type VenueInfo struct {
	SalesAreas []SalesArea `json:"salesAreas"`
}

type VenueInfoResponse struct {
	Data VenueInfo `json:"data"`
}

type Menu struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MenusResponse struct {
	Data []Menu `json:"data"`
}

type Portion struct {
	Options []struct {
		Value struct {
			Price struct {
				Value float64 `json:"value"`
			} `json:"price"`
		} `json:"value"`
	} `json:"options"`
}

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Options     struct {
		Portion Portion `json:"portion"`
		Linked  []struct {
			Name string `json:"name"`
		} `json:"linked"`
	} `json:"options"`
}

type ItemGroup struct {
	Name  string `json:"name"`
	Items []Item `json:"items"`
}

type Category struct {
	Name       string      `json:"name"`
	ItemGroups []ItemGroup `json:"itemGroups"`
}

type MenuDetail struct {
	Categories []Category `json:"categories"`
}

type MenuDetailResponse struct {
	Data MenuDetail `json:"data"`
}

package fetcher

type Address struct {
	town     string
	location Location
}

type Location struct {
	longitude float64
	latitute  float64
}

type Venue struct {
	id       int
	venueRef int
	name     string
	location Address
}

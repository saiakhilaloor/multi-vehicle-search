package models

type Listing struct {
	ID           string `json:"id"`
	Length       int    `json:"length"`
	Width        int    `json:"width"`
	LocationID   string `json:"location_id"`
	PriceInCents int    `json:"price_in_cents"`
}
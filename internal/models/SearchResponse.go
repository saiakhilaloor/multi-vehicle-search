package models

type SearchResult struct {
	LocationID        string   `json:"location_id"`
	ListingIDs        []string `json:"listing_ids"`
	TotalPriceInCents int      `json:"total_price_in_cents"`
}
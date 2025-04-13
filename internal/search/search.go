package search

import (
	"github.com/saiakhilaloor/multi-vehicle-search/internal/models"

	"math"
	"sort"
)

type listingCombo struct {
	listingIDs        []string
	totalPriceInCents int
}

func FindMatches(vehicles []models.VehicleRequest, listings []models.Listing) []models.SearchResult {
	locationMap := make(map[string][]models.Listing)
	for _, l := range listings {
		locationMap[l.LocationID] = append(locationMap[l.LocationID], l)
	}

	var results []models.SearchResult

	for locationID, listings := range locationMap {
		bestCombo := findCheapestCombo(locationID, listings, vehicles)
		if bestCombo != nil {
			results = append(results, models.SearchResult{
				LocationID:         locationID,
				ListingIDs:         bestCombo.listingIDs,
				TotalPriceInCents:  bestCombo.totalPriceInCents,
			})
		}
	}

	// Sort by price ascending
	sort.Slice(results, func(i, j int) bool {
		return results[i].TotalPriceInCents < results[j].TotalPriceInCents
	})

	return results
}

func findCheapestCombo(locationID string, listings []models.Listing, vehicles []models.VehicleRequest) *listingCombo {
	n := len(listings)
	best := &listingCombo{totalPriceInCents: math.MaxInt}
	found := false

	for mask := 1; mask < (1 << n); mask++ {
		var selected []models.Listing
		var selectedIDs []string
		totalPrice := 0

		for i := 0; i < n; i++ {
			if mask&(1<<i) > 0 {
				selected = append(selected, listings[i])
				selectedIDs = append(selectedIDs, listings[i].ID)
				totalPrice += listings[i].PriceInCents
			}
		}

		if canFitAllVehicles(selected, vehicles) {
			if totalPrice < best.totalPriceInCents {
				best = &listingCombo{
					listingIDs:        selectedIDs,
					totalPriceInCents: totalPrice,
				}
				found = true
			}
		}
	}

	if found {
		return best
	}
	return nil
}

func canFitAllVehicles(listings []models.Listing, vehicles []models.VehicleRequest) bool {

	type slot struct {
		length int
		width  int
	}
	var slots []slot
	for _, l := range listings {
		slots = append(slots, slot{length: l.Length, width: l.Width})
	}

	for _, v := range vehicles {
		needed := v.Quantity

		for i := range slots {
			fitInRow := slots[i].length / v.Length
			fitRows := slots[i].width / 10
			fitTotal := fitInRow * fitRows

			take := min(needed, fitTotal)
			needed -= take

			if needed == 0 {
				break
			}
		}

		if needed > 0 {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

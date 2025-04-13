package listings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/saiakhilaloor/multi-vehicle-search/internal/models"
)

func LoadListings(filepath string) ([]models.Listing, error) {
	fmt.Println("Loading listings from:", filepath)

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	var listings []models.Listing
	if err := json.Unmarshal(bytes, &listings); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil, err
	}

	fmt.Printf("Successfully loaded %d listings\n", len(listings))
	return listings, nil
}

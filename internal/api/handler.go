package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saiakhilaloor/multi-vehicle-search/internal/listings"
	"github.com/saiakhilaloor/multi-vehicle-search/internal/models"
	"github.com/saiakhilaloor/multi-vehicle-search/internal/search"
)

func SearchHandler(c *gin.Context) {
	var vehicleRequests []models.VehicleRequest
	if err := c.ShouldBindJSON(&vehicleRequests); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	listingsData, err := listings.LoadListings("./listings.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not load listings"})
		return
	}

	results := search.FindMatches(vehicleRequests, listingsData)
	c.JSON(http.StatusOK, results)
}

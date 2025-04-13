package main

import (
	"github.com/gin-gonic/gin"
	"github.com/saiakhilaloor/multi-vehicle-search/internal/api"
)

func main() {
	router := gin.Default()

	router.POST("/search", api.SearchHandler)

	router.Run(":8080")
}

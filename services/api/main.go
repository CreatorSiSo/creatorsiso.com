package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Season struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

var seasons = []Season{
	{ID: "4ß508340157386", Name: "Winter", Active: false},
	{ID: "95703463864ß51", Name: "Spring", Active: true},
	{ID: "864ß508s573463", Name: "Summer", Active: false},
	{ID: "16sß5974034864", Name: "Fall", Active: false},
}

func getSeasons(context *gin.Context) {
	context.JSON(http.StatusOK, seasons)
}

func getSeason(context *gin.Context) {
	id := context.Param("id")

	for _, season := range seasons {
		if id == season.ID {
			context.JSON(http.StatusOK, season)
			return
		}
	}

	context.JSON(http.StatusNotFound, gin.H{
		"message": "Could not find season with id: " + id,
	})
}

func addSeason(context *gin.Context) {
	var newSeason Season

	if err := context.BindJSON(&newSeason); err == nil {
		seasons = append(seasons, newSeason)
		context.JSON(http.StatusCreated, seasons)
	}
}

func main() {
	port, exists := os.LookupEnv("API_PORT")
	if !exists {
		fmt.Println("The API_PORT environment variable is not set! => Defaulting to 6333")
		port = "6333"
	}

	router := gin.Default()
	router.RedirectTrailingSlash = true

	router.GET("/")

	router.GET("/seasons", getSeasons)
	router.GET("/seasons/:id", getSeason)
	router.POST("/seasons", addSeason)

	router.Run("localhost:" + port)
}

// func get_time(response http.ResponseWriter, request *http.Request) {
// 	fmt.Fprintf(response, "Current Time: %s %s", time.Now(), "\n")
// }

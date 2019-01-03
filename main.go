package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	apikey := os.Getenv("DARK_SKY_API_KEY")
	if len(apikey) == 0 {
		log.Fatalln("DARK_SKY_API_KEY environment variable must be set.")
	}

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/current_weather/:lat/:long", func(c *gin.Context) {
		lat, err := strconv.ParseFloat(c.Params.ByName("lat"), 64)
		long, err := strconv.ParseFloat(c.Params.ByName("long"), 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Latitude and Longitude are required"})
			return
		}
		response, err := currentWeather(lat, long, apikey)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err.Error()})
			return
		}
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, gin.H{"weather": response})
	})

	err := endless.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
}

func currentWeather(lat, long float64, apikey string) (ForecastResponse, error) {
	client := NewDarkSkyAPI(apikey)

	request := ForecastRequest{}
	request.Latitude = lat
	request.Longitude = long
	request.Options = ForecastRequestOptions{
		Exclude: "minutely",
		Lang:    "en",
		Units:   "us",
	}

	return client.Forecast(request)
}

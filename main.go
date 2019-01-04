package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	apikey := os.Getenv("DARK_SKY_API_KEY")
	if len(apikey) == 0 {
		log.Fatalln("DARK_SKY_API_KEY environment variable must be set.")
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "https://github.com/leothelocust/dark-sky-weather-api")
	})
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.File("assets/favicon.ico")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	cacheDuration := time.Minute * 15
	cacheStore := persistence.NewInMemoryStore(cacheDuration)

	router.GET("/current_weather/:lat/:long", cache.CachePage(cacheStore, cacheDuration, func(c *gin.Context) {
		lat, err := strconv.ParseFloat(c.Params.ByName("lat"), 64)
		long, err := strconv.ParseFloat(c.Params.ByName("long"), 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Latitude and Longitude are required"})
			return
		}
		response, err := currentWeather(lat, long, apikey)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err.Error()})
			log.Printf("Error: %s\n", err)
			return
		}
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, gin.H{"weather": response})
	}))

	if os.Getenv("GIN_MODE") == "release" {
		log.Fatal(autotls.Run(router, "weather.l3vi.co"))
	} else {
		log.Fatal(router.Run(":80"))
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

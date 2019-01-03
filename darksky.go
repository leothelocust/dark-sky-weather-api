package main

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// DarkSky API endpoint
var (
	BaseURL = "https://api.darksky.net/forecast"
)

// DarkSky Api client
type DarkSky interface {
	Forecast(request ForecastRequest) (ForecastResponse, error)
}

type darkSky struct {
	APIKey string
	Client *http.Client
}

// NewDarkSkyAPI creates a new DarkSky client
func NewDarkSkyAPI(apiKey string) DarkSky {
	return &darkSky{apiKey, &http.Client{}}
}

// Forecast request a forecast
func (d *darkSky) Forecast(request ForecastRequest) (ForecastResponse, error) {
	response := ForecastResponse{}

	url := d.buildRequestURL(request)

	err := get(d.Client, url, &response)

	return response, err
}

func (d *darkSky) buildRequestURL(request ForecastRequest) string {
	url := fmt.Sprintf("%s/%s/%f,%f", BaseURL, d.APIKey, request.Latitude, request.Longitude)

	if request.Time > 0 {
		url = url + fmt.Sprintf(",%d", request.Time)
	}

	values, _ := query.Values(request.Options)
	queryString := values.Encode()

	if len(queryString) > 0 {
		url = url + "?" + queryString
	}

	return url
}

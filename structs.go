package main

// Timestamp is an int64 timestamp
type Timestamp int64

// ForecastRequest contains all available options for requesting a forecast
type ForecastRequest struct {
	Latitude  float64
	Longitude float64
	Time      Timestamp
	Options   ForecastRequestOptions
}

// ForecastRequestOptions are optional and passed as query parameters
type ForecastRequestOptions struct {
	Exclude string `url:"exclude,omitempty"`
	Extend  string `url:"extend,omitempty"`
	Lang    string `url:"lang,omitempty"`
	Units   string `url:"units,omitempty"`
}

// ForecastResponse is the response containing all requested properties
type ForecastResponse struct {
	Latitude  float64    `json:"latitude,omitempty"`
	Longitude float64    `json:"longitude,omitempty"`
	Timezone  string     `json:"timezone,omitempty"`
	Currently *DataPoint `json:"currently,omitempty"`
	Minutely  *DataBlock `json:"minutely,omitempty"`
	Hourly    *DataBlock `json:"hourly,omitempty"`
	Daily     *DataBlock `json:"daily,omitempty"`
	Alerts    []*Alert   `json:"alerts,omitempty"`
	Flags     *Flags     `json:"flags,omitempty"`
}

// DataPoint contains various properties, each representing the average (unless otherwise specified) of a particular weather phenomenon occurring during a period of time.
type DataPoint struct {
	ApparentTemperature     float64   `json:"apparentTemperature,omitempty"`
	ApparentTemperatureHigh float64   `json:"apparentTemperatureHigh,omitempty"`
	ApparentTemperatureLow  float64   `json:"apparentTemperatureLow,omitempty"`
	Humidity                float64   `json:"humidity,omitempty"`
	Icon                    string    `json:"icon"`
	MoonPhase               float64   `json:"moonPhase,omitempty"`
	Summary                 string    `json:"summary,omitempty"`
	SunriseTime             Timestamp `json:"sunriseTime,omitempty"`
	SunsetTime              Timestamp `json:"sunsetTime,omitempty"`
	Temperature             float64   `json:"temperature,omitempty"`
	TemperatureHigh         float64   `json:"temperatureHigh"`
	TemperatureLow          float64   `json:"temperatureLow"`
	TemperatureMax          float64   `json:"temperatureMax"`
	TemperatureMin          float64   `json:"temperatureMin"`
	Time                    Timestamp `json:"time,omitempty"`
	Visibility              float64   `json:"visibility,omitempty"`
	WindBearing             float64   `json:"windBearing"`
	WindGust                float64   `json:"windGust"`
	WindSpeed               float64   `json:"windSpeed"`
}

// DataBlock represents the various weather phenomena occurring over a period of time
type DataBlock struct {
	Summary string      `json:"summary,omitempty"`
	Icon    string      `json:"icon,omitempty"`
	Data    []DataPoint `json:"data,omitempty"`
}

// Alert contains objects representing the severe weather warnings issued for the requested location by a governmental authority
type Alert struct {
	Title       string    `json:"title,omitempty"`
	Severity    string    `json:"severity,omitempty"`
	Description string    `json:"description,omitempty"`
	Expires     Timestamp `json:"expires,omitempty"`
	Regions     []string  `json:"regions,omitempty"`
	Time        Timestamp `json:"time,omitempty"`
	URI         string    `json:"uri,omitempty"`
}

// Flags contains various metadata information related to the request
type Flags struct {
	DarkSkyUnavailable string   `json:"darksky-unavailable,omitempty"`
	Sources            []string `json:"sources,omitempty"`
	Units              string   `json:"units,omitempty"`
}

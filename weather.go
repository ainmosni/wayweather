package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const (
	// WeatherBase is the base URL for the wttr.in service.
	WeatherBase = "https://wttr.in/"
)

// getWeather gets the weather and returns the JSON body.
func getWeatherData(location string) (WeatherData, error) {
	u, err := url.Parse(WeatherBase)
	if err != nil {
		return WeatherData{}, fmt.Errorf(
			"couldn't parse base URL %s: %v", WeatherBase, err,
		)
	}

	if location != "" {
		u.Path = path.Join(u.Path, location)
	}

	q := u.Query()
	q.Set("format", "j1")
	u.RawQuery = q.Encode()

	r, err := http.Get(u.String())
	if err != nil {
		return WeatherData{}, fmt.Errorf("couldn't get %s: %v", u.String(), err)
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return WeatherData{}, fmt.Errorf(
			"couldn't read body from %s: %v", u.String(), err,
		)
	}
	return parseWeatherJSON(b)
}

// getSymbol takes a weather code and returns the proper emoji.
func getWeatherSymbol(weatherCode int) (symbol string) {
	switch weatherCode {
	case 113:
		return "🌣"
	case 116:
		return "🌤"
	case 119, 122:
		return "☁"
	case 143, 248, 260:
		return "🌫"
	case 176, 263, 281, 293, 296, 299, 302, 305, 308, 311, 353, 356, 359, 362:
		return "🌦"
	case 179, 182, 185, 227, 230, 284, 314, 317, 320, 323, 326, 329, 332, 335, 338, 350, 365, 368, 371, 374, 377:
		return "❄"
	case 200, 386, 389, 392, 395:
		return "🌩"
	default:
		return "⚠️"
	}
}

// Takes the JSON body from wttr.in and returns the condition, temperature, and region
func parseWeatherJSON(body []byte) (weatherData WeatherData, err error) {
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return
	}

	return
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type output struct {
	Text    string `json:"text,omitempty"`
	Alt     string `json:"alt,omitempty"`
	Tooltip string `json:"tooltip,omitempty"`
}

func printError(err error) bool {
	if err == nil {
		return false
	}

	o, err := json.Marshal(output{
		Text:    "Error",
		Alt:     "Error querying data.",
		Tooltip: err.Error(),
	})
	if err != nil {
		log.Fatalf("Couldn't marshall error: %v", err)
	}
	fmt.Printf("%s\n", string(o))
	return true
}
func printReport(wd WeatherData, err error) {
	if printError(err) {
		return
	}

	code, err := strconv.Atoi(wd.CurrentCondition[0].WeatherCode)
	if err != nil {
		code = 666
	}

	symbol := getWeatherSymbol(code)
	region := wd.NearestArea[0].Region[0]["value"]
	temp := wd.CurrentCondition[0].TempC

	o, err := json.Marshal(output{
		Text:    fmt.Sprintf("%s° %s", temp, symbol),
		Alt:     region,
		Tooltip: fmt.Sprintf("%s: %s° %s", region, temp, WeatherCodeToCondition[code]),
	})

	if printError(err) {
		return
	}

	fmt.Printf("%s\n", string(o))

}

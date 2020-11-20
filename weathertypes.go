package main

// WeatherData represents the main weather data structure, at least
// the parts that we're interested in.
type WeatherData struct{
	CurrentCondition []Condition `json:"current_condition,omitempty"`
	NearestArea      []Area        `json:"nearest_area,omitempty"`
}

// Condition are the weather conditions. This API encodes everything as a
// string. :(
type Condition struct{
	TempC       string `json:"temp_C,omitempty"`
	WeatherCode string `json:"weatherCode,omitempty"`
}


// Area is a stucture that describes an area.
type Area struct {
	Region []map[string]string `json:"region,omitempty"`
}

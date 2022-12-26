package types

type Response struct {
	Code int `json:"cod"`
}

type Weather struct {
	Main []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Temperature struct {
		TemperatureF float64 `json:"temp"`
	} `json:"main"`
}

type WeatherResponse struct {
	Temperature float64 `json:"temperature"`
	Description string  `json:"description"`
}

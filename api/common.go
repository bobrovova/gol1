package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/bobrovova/go-weather/configs"
	"github.com/bobrovova/go-weather/pkg/types"
)

func SendResponse(w *http.ResponseWriter, data any) error {
	jsonResponse, err := json.Marshal(data)

	if err != nil {
		return err
	}

	(*w).Header().Add("Content-Type", "application/json")
	(*w).Write(jsonResponse)

	return nil
}

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data types.Response

	json.Unmarshal(body, &data)

	if data.Code != 200 {
		return nil, errors.New("Fetch problem: " + url)
	}

	return body, nil
}

func FetchCurrentTemperature(city string) (*types.Weather, error) {
	requestValues := url.Values{
		"q":     {city},
		"appid": {configs.APIKey},
	}
	url := url.URL{
		Scheme:   "https",
		Host:     "api.openweathermap.org",
		Path:     "/data/2.5/weather",
		RawQuery: requestValues.Encode(),
	}
	body, err := Fetch(url.String())
	if err != nil {
		return nil, err
	}

	var weather types.Weather
	json.Unmarshal(body, &weather)

	return &weather, nil
}

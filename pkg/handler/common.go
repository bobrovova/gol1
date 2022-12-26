package handler

import (
	"net/http"
	"net/url"

	"github.com/bobrovova/go-weather/api"
	_ "github.com/bobrovova/go-weather/docs"
	"github.com/bobrovova/go-weather/pkg/types"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

func InitRoutes(logger *logrus.Logger) chi.Router {
	r := chi.NewRouter()
	r.Get("/", CurrentTemperatureHandler(logger))
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8881/swagger/swagger.json"),
	))

	return r
}

// @Summary CurrentTemperature
// @Tags Temperature
// @Accept json
// @Produce json
func CurrentTemperatureHandler(logger *logrus.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		values, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			logger.Error(err)
			return
		}

		if values["city"] != nil {
			city := values["city"][0]

			weather, err := api.FetchCurrentTemperature(city)
			if err != nil {
				logger.Error(err)
				return
			}

			err = api.SendResponse(&w, types.WeatherResponse{
				Temperature: weather.Temperature.TemperatureF - 273.15,
				Description: weather.Main[0].Description,
			})

			if err != nil {
				logger.Error(err)
			}
		}
	}
}

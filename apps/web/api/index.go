// @title Weather-Flex API
// @version 1.0
// @description This is a Go-powered weather proxy for OpenWeatherMap.
// @contact.name Anna Shulik
// @contact.url https://github.com/tmnts
// @host weather-flex.vercel.app
// @BasePath /api
package handler // Vercel NEEDS this to be "handler"

import (
	httpSwagger "://github.com"
	"io"
	"net/http"
	"os"
	"strings"
	_ "weather-api/docs"

	httpSwagger "github.com/swaggo/swag/cmd/swag@latest"
)

// Handler godoc
// @Summary Get current weather
// @Description Fetch real-time weather data for Khimki via OpenWeatherMap
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /weather [get]
func Handler(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	// Make sure this URL is 100% correct with your lat/lon
	apiUrl := "https://api.openweathermap.org/data/2.5/weather?q=moscow&units=metric&appid=" + apiKey

	if strings.Contains(r.URL.Path, "swagger") {
		httpSwagger.WrapHandler(w, r)
		return // Exit here so we don't run the weather code
	}

	resp, err := http.Get(apiUrl)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	// Copy the data to the response
	io.Copy(w, resp.Body)
}

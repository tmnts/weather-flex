package handler // Vercel NEEDS this to be "handler"

import (
	"io"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	// Make sure this URL is 100% correct with your lat/lon
	apiUrl := "https://api.openweathermap.org/data/2.5/weather?q=moscow&units=metric&appid=" + apiKey

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

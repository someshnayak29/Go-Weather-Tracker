package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

// Get api key from .apiConfig file
func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err
	}

	// we are storing it in json and not using directly, bcoz with json data we can't perform operations in golang, therefore use struct

	var c apiConfigData

	// unmarshall because filename in .apiConfig file is in .json, therefore
	// unmarshall json into struct

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}

	return c, nil
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Weatheria!!!\n"))
}

func query(city string) (weatherData, error) {

	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return weatherData{}, err
	}

	// Query parameters are key-value pairs appended to the end of a URL after a question mark ?. They are separated by ampersands &.
	// e.g. http://localhost:8080/info?name=John&age=30

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + city)

	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	return d, nil

}

func main() {

	http.HandleFunc("/hello", hello)

	// splits the r.URL.Path string into parts separated by "/" and assigns the third part (index 2) to the variable city.
	// "/weather/forecast/new-york" => "", "/weather", "forecast/new-york" split into 3 parts and access [2] index i.e. forecast/new-york

	http.HandleFunc("/weather/",

		func(w http.ResponseWriter, r *http.Request) {
			city := strings.SplitN(r.URL.Path, "/", 3)[2]
			data, err := query(city)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(data)
		})

	http.ListenAndServe(":8080", nil)
}

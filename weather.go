package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Run weather forecast")
	// call API
	weather, err := LoadWeather()
	// out to console
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("Weather in Mayamy: ", weather)
	}
}

var apiURL string = "http://api.weatherapi.com/v1/current.json?key=ff73ea28892b4a7bb40175852210110&q=Mayami"

func LoadWeather() (string, error) {
	rawRespone, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer rawRespone.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rawRespone.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

package config

import "os"

var WeatherToken string

func GetWeatherToken() string {
	if WeatherToken == "" {
		WeatherToken = os.Getenv("WEATHER_TOKEN")
	}
	return WeatherToken
}

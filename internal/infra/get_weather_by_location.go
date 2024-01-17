package infra

import (
	"encoding/json"
	"io"
	"net/http"

	configs "github.com/ItaloG/go-weather-api/config"
)

type WeatherApiResponse struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            int     `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb int     `json:"pressure_mb"`
		PressureIn int     `json:"pressure_in"`
		PrecipMm   int     `json:"precip_mm"`
		PrecipIn   int     `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      int     `json:"vis_km"`
		VisMiles   int     `json:"vis_miles"`
		Uv         int     `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
}

type WeatherTemperature struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func GetWeatherByLocation(location string) (WeatherTemperature, error) {
	config, err := configs.LoadConfig()
	if err != nil {
		return WeatherTemperature{}, err
	}
	resp, err := http.Get("https://api.weatherapi.com/v1/current.json?q=" + location + "&key=" + config.WeatherApiKey)
	if err != nil {
		return WeatherTemperature{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WeatherTemperature{}, err
	}

	var weatherResponse WeatherApiResponse
	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		return WeatherTemperature{}, err
	}

	tempK := float64(weatherResponse.Current.TempC) + 273

	return WeatherTemperature{
		TempC: float64(weatherResponse.Current.TempC),
		TempF: float64(weatherResponse.Current.TempF),
		TempK: tempK,
	}, nil
}

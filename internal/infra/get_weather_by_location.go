package infra

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
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
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree float64 `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   float64 `json:"humidity"`
		Cloud      float64 `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
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
	formattedLocation := strings.Replace(location, " ", "_", -1)

	resp, err := http.Get("http://api.weatherapi.com/v1/current.json?q=" + formattedLocation + "&key=ffc0bf9a91504a27abe01823241701")
	if err != nil {
		return WeatherTemperature{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return WeatherTemperature{}, errors.New("unable to find weather")
	}
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

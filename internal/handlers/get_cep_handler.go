package handlers

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/ItaloG/go-weather-api/config"
	"github.com/ItaloG/go-weather-api/internal/infra"
	"github.com/go-chi/chi/v5"
)

func GetCepHandler(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")

	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("you must pass a zipcode")
		return
	}

	isValid, err := regexp.MatchString(`\d{8}`, cep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("unable to process request")
		return
	}
	if !isValid {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("invalid zipcode")
		return
	}

	location, err := infra.GetCepFromViaCep(cep, "http://viacep.com.br/ws/")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("can not found zipcode")
		return
	}
	weather, err := infra.GetWeatherByLocation(location, "http://api.weatherapi.com/v1/current.json?q=", config.GetWeatherToken())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("can not found weather")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weather)
}

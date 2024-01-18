package main

import (
	"net/http"

	configs "github.com/ItaloG/go-weather-api/config"
	"github.com/ItaloG/go-weather-api/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Get("/{cep}", handlers.GetCepHandler)

	http.ListenAndServe(":8000", r)
}

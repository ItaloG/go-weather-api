package main

import (
	"net/http"

	"github.com/ItaloG/go-weather-api/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/{cep}", handlers.GetCepHandler)

	http.ListenAndServe(":8080", r)
}

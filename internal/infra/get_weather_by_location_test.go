package infra_test

import (
	"net/http/httptest"
	"testing"

	"github.com/ItaloG/go-weather-api/internal/infra"
	"github.com/stretchr/testify/assert"
)

func TestGivenAValidLocation_ShouldReturnTheWeatherTemp(t *testing.T) {
	mockServer := httptest.NewServer(&MockHTTPServer{
		StatusCode:   200,
		ResponseBody: `{"current": {"temp_c": 20.0, "temp_f": 69.9}}`,
	})
	defer mockServer.Close()

	validLocation := "São Paulo"

	weather, err := infra.GetWeatherByLocation(validLocation, mockServer.URL+"/", "valid_token")
	assert.NotEmpty(t, weather.TempC)
	assert.NotEmpty(t, weather.TempF)
	assert.NotEmpty(t, weather.TempK)
	assert.Nil(t, err)
}

func TestGivenAnInvalidLocation_ShouldReturnAnError(t *testing.T) {
	mockServer := httptest.NewServer(&MockHTTPServer{
		StatusCode:   400,
		ResponseBody: `{"message": "Server error"}`,
	})
	defer mockServer.Close()

	invalidLocation := "invalid_location"

	weather, err := infra.GetWeatherByLocation(invalidLocation, mockServer.URL+"/", "valid_token")
	assert.Empty(t, weather)
	assert.Equal(t, err.Error(), "unable to find weather")
}

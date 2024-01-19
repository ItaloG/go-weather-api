package infra_test

import (
	"testing"

	"github.com/ItaloG/go-weather-api/internal/infra"
	"github.com/stretchr/testify/assert"
)

func TestGivenAValidLocation_ShouldReturnTheWeatherTemp(t *testing.T) {

	validLocation := "São Paulo"

	weather, err := infra.GetWeatherByLocation(validLocation)
	assert.NotEmpty(t, weather.TempC)
	assert.NotEmpty(t, weather.TempF)
	assert.NotEmpty(t, weather.TempK)
	assert.Nil(t, err)
}

func TestGivenAnInvalidLocation_ShouldReturnAnError(t *testing.T) {

	invalidLocation := "invalid_location"

	weather, err := infra.GetWeatherByLocation(invalidLocation)
	assert.Empty(t, weather)
	assert.Equal(t, err.Error(), "unable to find weather")
}

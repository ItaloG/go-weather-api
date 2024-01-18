package infra_test

import (
	"testing"

	"github.com/ItaloG/go-weather-api/internal/infra"
	"github.com/stretchr/testify/assert"
)

func TestGivenAValidCep_ShouldReturnTheLocalidade(t *testing.T) {

	validCep := "01001000" // Sé - São Paulo

	location, err := infra.GetCepFromViaCep(validCep)
	assert.Equal(t, location, "São Paulo")
	assert.Nil(t, err)
}

func TestGivenAnInvalidCep_ShouldReturnAnError(t *testing.T) {

	invalidCep := "00000000"

	location, err := infra.GetCepFromViaCep(invalidCep)
	assert.Equal(t, location, "")
	assert.Equal(t, err.Error(), "unable to find cep")
}

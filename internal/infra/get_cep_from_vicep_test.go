package infra_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ItaloG/go-weather-api/internal/infra"
	"github.com/stretchr/testify/assert"
)

type MockHTTPServer struct {
	StatusCode   int
	ResponseBody string
}

func (m *MockHTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(m.StatusCode)
	w.Write([]byte(m.ResponseBody))
}

func TestGivenAValidCep_ShouldReturnTheLocalidade(t *testing.T) {
	mockServer := httptest.NewServer(&MockHTTPServer{
		StatusCode:   200,
		ResponseBody: `{"localidade": "São Paulo"}`,
	})
	defer mockServer.Close()

	validCep := "01001000"

	location, err := infra.GetCepFromViaCep(validCep, mockServer.URL+"/")
	assert.Equal(t, location, "São Paulo")
	assert.Nil(t, err)
}

func TestGivenAnInvalidCep_ShouldReturnAnError(t *testing.T) {
	mockServer := httptest.NewServer(&MockHTTPServer{
		StatusCode:   200,
		ResponseBody: `{"erro": true}`,
	})
	defer mockServer.Close()

	invalidCep := "00000000"

	location, err := infra.GetCepFromViaCep(invalidCep, mockServer.URL+"/")
	assert.Equal(t, location, "")
	assert.Equal(t, err.Error(), "unable to find cep")
}

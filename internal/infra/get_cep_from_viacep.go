package infra

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
	Erro        bool   `json:"erro"`
}

func GetCepFromViaCep(cep string, url string) (string, error) {
	resp, err := http.Get(url + cep + "/json/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", errors.New("unable to find cep")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var c ViaCEP
	err = json.Unmarshal(body, &c)
	if err != nil {
		return "", err
	}
	if c.Erro {
		return "", errors.New("unable to find cep")
	}

	return c.Localidade, nil
}

package services

import (
	"encoding/json"
	"io"
	"linknau-test/models"
	"net/http"
)

func FetchDataFromRemote(url string) ([]models.Data, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []models.Data{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []models.Data{}, err
	}

	var data []models.Data
	if err := json.Unmarshal(body, &data); err != nil {
		return []models.Data{}, err
	}

	return data, nil
}

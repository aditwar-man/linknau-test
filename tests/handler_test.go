package tests

import (
	"encoding/json"
	"linknau-test/models"
	"linknau-test/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchData(t *testing.T) {
	mockData := []models.Data{
		{ID: 1, UserId: 101, Title: "Test1", Completed: false},
		{ID: 2, UserId: 102, Title: "Test2", Completed: true},
	}
	mockDataBytes, _ := json.Marshal(mockData)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(mockDataBytes)
	}))
	defer server.Close()

	url := server.URL
	data, err := services.FetchDataFromRemote(url)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	t.Logf("Fetched Data: %+v", data)

	if len(data) != len(mockData) {
		t.Errorf("Expected %v elements, got %v", len(mockData), len(data))
	}

	for i := range data {
		if data[i].ID != mockData[i].ID || data[i].UserId != mockData[i].UserId || data[i].Title != mockData[i].Title || data[i].Completed != mockData[i].Completed {
			t.Errorf("Expected %v, got %v", mockData[i], data[i])
		}
	}
}

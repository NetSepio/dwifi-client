package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"chicken/wifi"
)

func GetWiFiDataFromHTTP() ([]wifi.WiFiData, error) {
	url := "https://dev.gateway.erebrus.io/api/v1.0/nodedwifi/all"

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from API: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var apiResponse struct {
		Data []wifi.WiFiData `json:"data"`
	}
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return apiResponse.Data, nil
}

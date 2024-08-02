package weather

import (
	"encoding/json"
	"fmt"
)

const BaseURL = "https://api.openweathermap.org"

type Conditions struct {
	Summary string `json:"summary"`
}

type OWMResponse struct {
	Weather []struct {
		Main string
	}
}

func ParseResponse(data []byte) (Conditions, error) {
	var resp OWMResponse
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return Conditions{}, fmt.Errorf("invalid API response %s: %w", data, err)
	}
	if len(resp.Weather) < 1 {
		return Conditions{}, fmt.Errorf("API response has an emtpy weather field: %s", resp)
	}
	conditions := Conditions{
		Summary: resp.Weather[0].Main,
	}
	return conditions, nil
}

func FormatUrl(baseURL, location, key string) string {
	return fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s", baseURL, location, key)
}

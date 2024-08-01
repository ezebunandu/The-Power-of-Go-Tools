package weather

import (
	"encoding/json"
	"fmt"
)

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
	conditions := Conditions{
		Summary: resp.Weather[0].Main,
	}
	return conditions, nil
}

package weather_test

import (
	"os"
	"testing"

	"github.com/ezebunandu/weather"
	"github.com/google/go-cmp/cmp"
)

func TestParseResponse_CorrectlyParsesJSONData(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/weather.json")
	if err != nil {
		t.Fatal(err)
	}
	want := weather.Conditions{
		Summary: "Clouds",
	}
	got, err := weather.ParseResponse(data)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestParseResponse_ReturnsErrorForGivenEmptyData(t *testing.T) {
	t.Parallel()
	_, err := weather.ParseResponse([]byte{})
	if err == nil {
		t.Fatal(err)
	}
}

func TestParseResponse_ReturnsErrorGivenInvalidJSON(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/invalid-weather.json")
	if err != nil {
		t.Fatal(err)
	}
	_, err = weather.ParseResponse(data)
	if err == nil {
		t.Error("want error when parsing invalid response, got nil")
	}
}

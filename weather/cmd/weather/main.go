package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ezebunandu/weather"
)

func main() {
	key := os.Getenv("OPENWEATHERMAP_API_KEY")
	if key == "" {
		fmt.Fprintln(os.Stderr, "Please set the environment variable OPENWEATHER_API_KEY.")
		os.Exit(1)
	}
	location := "Calgary,CA"
	URL := weather.FormatUrl(weather.BaseURL, location, key)
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintln(os.Stderr, "unexpected response status", resp.Status)
		os.Exit(1)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	conditions, err := weather.ParseResponse(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(conditions)
}

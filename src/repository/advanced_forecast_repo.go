package repository

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchAdvancedForecastHTML(beach string) (string, error) {
	url := fmt.Sprintf("https://www.surf-forecast.com/breaks/%s/forecasts/latest/six_day", beach)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error fetching URL: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}
	return string(body), nil
}

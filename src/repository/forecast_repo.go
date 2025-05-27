package repository

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchForecastHTML(beach string) (string, error) {
	url := fmt.Sprintf("https://www.surf-forecast.com/breaks/%s/forecasts/widget/a", beach)
	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error fetching URL: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	return string(body), nil
}

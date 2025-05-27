package service

import (
	"fmt"
	"strings"
	"time"

	"seaventures/src/models"
	"seaventures/src/repository"

	"github.com/PuerkitoBio/goquery"
)

func get(arr []string, i int) string {
	if i < len(arr) {
		return arr[i]
	}
	return ""
}

func GetForecast(beach string) (*models.MultiDayResult, error) {
	html, err := repository.FetchForecastHTML(beach)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	grid := doc.Find("div.grid")
	if grid.Length() == 0 {
		return nil, fmt.Errorf("forecast grid not found")
	}

	extractRow := func(class string) []string {
		var out []string
		grid.Find("div." + class).Find("div.cell").Each(func(i int, s *goquery.Selection) {
			out = append(out, strings.TrimSpace(s.Text()))
		})
		return out
	}

	times := extractRow("time-row.row")
	ratings := extractRow("rating-row.row")
	heights := []string{}
	grid.Find("div.wave-row.row div.cell").Each(func(i int, s *goquery.Selection) {
		height := s.Find("span.height").Text()
		heights = append(heights, strings.TrimSpace(height))
	})
	periods := extractRow("period-row.row")
	speeds := []string{}
	grid.Find("div.wind-row.row div.cell").Each(func(i int, s *goquery.Selection) {
		val := s.Find("div.wind-icon text").Text()
		if val == "" {
			val = s.Find("div.wind-icon svg text").Text()
		}
		speeds = append(speeds, strings.TrimSpace(val))
	})
	states := extractRow("row.wind-state-row")
	weather := []string{}
	grid.Find("div.row.weather-state-row div.cell").Each(func(i int, s *goquery.Selection) {
		img := s.Find("img.weather-icon")
		if img.Length() > 0 {
			alt, _ := img.Attr("alt")
			weather = append(weather, alt)
		} else {
			weather = append(weather, "")
		}
	})
	airtemps := []string{}
	grid.Find("div.row.temperature-row div.cell").Each(func(i int, s *goquery.Selection) {
		temp := s.Find("span.temp").Text()
		if temp != "" {
			airtemps = append(airtemps, temp+"°C")
		} else {
			airtemps = append(airtemps, "")
		}
	})

	// Water temp
	waterTemp := ""
	grid.Find("div.row.current-weather-row span.water-temp span.temp").Each(func(i int, s *goquery.Selection) {
		waterTemp = s.Text() + "°C"
	})

	forecast := []models.ForecastEntry{}
	n := len(times)
	for i := 0; i < n; i++ {
		entry := models.ForecastEntry{
			Time:           get(times, i),
			Rating:         get(ratings, i),
			Height:         get(heights, i),
			Period:         get(periods, i),
			SpeedDirection: get(speeds, i),
			State:          get(states, i),
			Weather:        get(weather, i),
			AirTemp:        get(airtemps, i),
		}
		forecast = append(forecast, entry)
	}

	entriesPerDay := 8
	numDays := (len(forecast) + entriesPerDay - 1) / entriesPerDay

	days := []models.DayForecast{}
	startDate := time.Now().Format("2006-01-02")
	for day := 0; day < numDays; day++ {
		start := day * entriesPerDay
		end := start + entriesPerDay
		if end > len(forecast) {
			end = len(forecast)
		}
		dayForecast := forecast[start:end]

		date := ""
		if t, err := time.Parse("2006-01-02", startDate); err == nil {
			date = t.AddDate(0, 0, day).Format("2006-01-02")
		} else {
			date = fmt.Sprintf("day%d", day+1)
		}

		days = append(days, models.DayForecast{
			Date:     date,
			Forecast: dayForecast,
		})
	}

	result := &models.MultiDayResult{
		WaterTemp: waterTemp,
		Days:      days,
	}

	return result, nil
}

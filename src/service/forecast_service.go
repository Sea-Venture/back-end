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

type SurfTableStructured struct {
	Dates       []string   `json:"dates"`
	TimeRanges  []string   `json:"timeRanges"`
	SwellHeight []string   `json:"swellHeight,omitempty"`
	Period      []string   `json:"period,omitempty"`
	Energy      []string   `json:"energy,omitempty"`
	Wind        []string   `json:"wind,omitempty"`
	WindState   []string   `json:"windState,omitempty"`
	HighTide    []string   `json:"highTide,omitempty"`
	LowTide     []string   `json:"lowTide,omitempty"`
	Other       [][]string `json:"other,omitempty"`
	TfootRows   [][]string `json:"tfootRows,omitempty"`
}

func GetSurfTable(beach string) (*SurfTableStructured, error) {
	html, err := repository.FetchForecastHTML(beach)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	table := doc.Find("table.js-forecast-table-content")
	if table.Length() == 0 {
		return nil, fmt.Errorf("forecast table not found")
	}

	// Extract headers
	headers := []string{}
	table.Find("thead th").Each(func(i int, s *goquery.Selection) {
		headers = append(headers, strings.TrimSpace(s.Text()))
	})

	// Extract tbody rows
	tbodyRows := [][]string{}
	table.Find("tbody tr").Each(func(i int, tr *goquery.Selection) {
		row := []string{}
		tr.Find("td").Each(func(j int, td *goquery.Selection) {
			row = append(row, strings.TrimSpace(td.Text()))
		})
		tbodyRows = append(tbodyRows, row)
	})

	// Extract tfoot rows
	tfootRows := [][]string{}
	table.Find("tfoot tr").Each(func(i int, tr *goquery.Selection) {
		row := []string{}
		tr.Find("td").Each(func(j int, td *goquery.Selection) {
			row = append(row, strings.TrimSpace(td.Text()))
		})
		tfootRows = append(tfootRows, row)
	})

	// Remove unwanted rows from tbodyRows
	filteredTbody := [][]string{}
	for _, row := range tbodyRows {
		if len(row) == 1 && (row[0] == "Weather\nSurf Details\nLocal Wavefinder\nGlobal Wavefinder" ||
			row[0] == "Weather Surf Details Local Wavefinder Global Wavefinder") {
			continue
		}
		filteredTbody = append(filteredTbody, row)
	}

	// Structure the JSON output
	structured := SurfTableStructured{
		TfootRows: tfootRows,
	}

	if len(filteredTbody) >= 2 {
		structured.Dates = filteredTbody[0]
		structured.TimeRanges = filteredTbody[1]
		dataRows := filteredTbody[2:]

		if len(dataRows) > 0 {
			structured.SwellHeight = dataRows[0]
		}
		if len(dataRows) > 1 {
			structured.Period = dataRows[1]
		}
		if len(dataRows) > 2 {
			structured.Energy = dataRows[2]
		}
		if len(dataRows) > 3 {
			structured.Wind = dataRows[3]
		}
		if len(dataRows) > 4 {
			structured.WindState = dataRows[4]
		}
		if len(dataRows) > 5 {
			structured.HighTide = dataRows[5]
		}
		if len(dataRows) > 6 {
			structured.LowTide = dataRows[6]
		}
		if len(dataRows) > 7 {
			structured.Other = dataRows[7:]
		}
	}

	return &structured, nil
}

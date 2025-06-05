package service

import (
	"fmt"
	"strings"

	"seaventures/src/models"
	"seaventures/src/repository"

	"github.com/PuerkitoBio/goquery"
)

func GetAdvancedForecast(beach string) (*models.AdvancedForecast, error) {
	html, err := repository.FetchAdvancedForecastHTML(beach)
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
	adv := models.AdvancedForecast{
		TfootRows: tfootRows,
	}

	if len(filteredTbody) >= 2 {
		adv.Dates = filteredTbody[0]
		adv.TimeRanges = filteredTbody[1]
		dataRows := filteredTbody[2:]

		if len(dataRows) > 0 {
			adv.SwellHeight = dataRows[0]
		}
		if len(dataRows) > 1 {
			adv.Period = dataRows[1]
		}
		if len(dataRows) > 2 {
			adv.Energy = dataRows[2]
		}
		if len(dataRows) > 3 {
			adv.Wind = dataRows[3]
		}
		if len(dataRows) > 4 {
			adv.WindState = dataRows[4]
		}
		if len(dataRows) > 5 {
			adv.HighTide = dataRows[5]
		}
		if len(dataRows) > 6 {
			adv.LowTide = dataRows[6]
		}
		if len(dataRows) > 7 {
			adv.Other = dataRows[7:]
		}
	}

	return &adv, nil
}

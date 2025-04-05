package service

import (
    "strings"

    "github.com/PuerkitoBio/goquery"
	"seaventures/src/models"
    "seaventures/src/repository"
)

func GetForecast(city string) ([]models.Forecast, error) {
    // Fetch raw HTML
    html, err := repository.FetchForecastHTML(city)
    if err != nil {
        return nil, err
    }

    // Parse the HTML string into a goquery document
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
    if err != nil {
        return nil, err
    }

    var forecasts []models.Forecast

    doc.Find("table").Each(func(tableIndex int, table *goquery.Selection) {
        headers := []string{}

        table.Find("thead tr th").Each(func(i int, th *goquery.Selection) {
            headers = append(headers, strings.TrimSpace(th.Text()))
        })

        table.Find("tbody tr").Each(func(rowIndex int, row *goquery.Selection) {
            var rowData []string
            row.Find("td").Each(func(i int, cell *goquery.Selection) {
                rowData = append(rowData, strings.TrimSpace(cell.Text()))
            })

            if len(rowData) == len(headers) {
                forecast := models.Forecast{
                    WaveType:    rowData[0],
                    TimeAndDate: rowData[1],
                    WaveHeight:  rowData[2],
                }
                forecasts = append(forecasts, forecast)
            }
        })
    })

    return forecasts, nil
}
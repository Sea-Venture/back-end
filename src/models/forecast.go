package models

type ForecastEntry struct {
	Time           string `json:"time"`
	Rating         string `json:"rating"`
	Height         string `json:"height"`
	Period         string `json:"period"`
	SpeedDirection string `json:"speed_direction"`
	State          string `json:"state"`
	Weather        string `json:"weather"`
	AirTemp        string `json:"air_temp"`
}

type DayForecast struct {
	Date     string          `json:"date"`
	Forecast []ForecastEntry `json:"forecast"`
}

type MultiDayResult struct {
	WaterTemp string        `json:"water_temp"`
	Days      []DayForecast `json:"days"`
}

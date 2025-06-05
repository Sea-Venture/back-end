package models

type AdvancedForecast struct {
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

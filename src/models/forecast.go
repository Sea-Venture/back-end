package models

type Forecast struct {
    WaveType    string `json:"Wave Type"`
    TimeAndDate string `json:"Time (+0530) & Date"`
    WaveHeight  string `json:"Wave Height & Period"`
}


package controller

import (
    "encoding/json"
    "net/http"

    "seaventures/src/service"
)

func GetForecastHandler(w http.ResponseWriter, r *http.Request) {
    city := r.URL.Query().Get("city")
    if city == "" {
        http.Error(w, "City parameter is required", http.StatusBadRequest)
        return
    }

    forecasts, err := service.GetForecast(city)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(forecasts)
}
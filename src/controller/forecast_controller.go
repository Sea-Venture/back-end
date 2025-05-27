package controller

import (
	"encoding/json"
	"net/http"

	"seaventures/src/service"
)

func GetForecastHandler(w http.ResponseWriter, r *http.Request) {
	beach := r.URL.Query().Get("beach")
	if beach == "" {
		http.Error(w, "Beach parameter is required", http.StatusBadRequest)
		return
	}

	result, err := service.GetForecast(beach)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

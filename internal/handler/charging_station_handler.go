package handler

import (
	"encoding/json"
	"nearest-charging-stations/internal/repository"
	"net/http"
	"strconv"
)

type ChargingStationHandler struct {
	repo *repository.ChargingStationRepository
}

func NewChargingStationHandler(repo *repository.ChargingStationRepository) *ChargingStationHandler {
	return &ChargingStationHandler{
		repo: repo,
	}
}

func (h *ChargingStationHandler) FindNearestStations(
	w http.ResponseWriter,
	r *http.Request,
) {
	latStr := r.URL.Query().Get("lat")
	lngStr := r.URL.Query().Get("lng")
	radiusStr := r.URL.Query().Get("radius")

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		http.Error(w, "Invalid latitude", http.StatusBadRequest)
		return
	}

	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		http.Error(w, "Invalid longitude", http.StatusBadRequest)
		return
	}

	radius, err := strconv.ParseFloat(radiusStr, 64)
	if err != nil {
		http.Error(w, "Invalid radius", http.StatusBadRequest)
		return
	}

	stations, err := h.repo.FindNearestStations(
		r.Context(),
		lat,
		lng,
		radius,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(stations)
}

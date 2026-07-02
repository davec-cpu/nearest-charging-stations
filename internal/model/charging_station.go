package model

type ChargingStation struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Longtitude float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
	DistanceM  float64 `json:"distance_m"`
}

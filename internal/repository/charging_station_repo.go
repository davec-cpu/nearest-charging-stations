package repository

import (
	"context"
	"log"
	"nearest-charging-stations/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ChargingStationRepository struct {
	db *pgxpool.Pool
}

func NewChargingStationRepository(db *pgxpool.Pool) *ChargingStationRepository {
	return &ChargingStationRepository{db: db}
}

func (r *ChargingStationRepository) FindNearestStations(
	ctx context.Context,
	lat float64,
	lng float64,
	radius float64,
) ([]model.ChargingStation, error) {
	query := `
		SELECT
		id, 
		name,
		ST_X(location) AS longitude,
		ST_Y(location) AS latitude, 
		ST_Distance(
			location::geography,
			ST_SetSRID(ST_MakePoint($1, $2), 4326)::geography
		) AS distance_m
		FROM charging_stations
		WHERE ST_DWithin(
			location::geography,
			ST_SetSRID(ST_MakePoint($1, $2), 4326)::geography,
			$3
		)
		ORDER BY distance_m
		LIMIT 10;
	`
	log.Printf("executing...")
	rows, err := r.db.Query(ctx, query, lng, lat, radius)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []model.ChargingStation

	for rows.Next() {
		var cs model.ChargingStation

		err := rows.Scan(
			&cs.ID,
			&cs.Name,
			&cs.Longtitude,
			&cs.Latitude,
			&cs.DistanceM,
		)

		if err != nil {
			return nil, err
		}
		result = append(result, cs)
	}
	log.Printf("Found %d stations", len(result))
	return result, nil
}

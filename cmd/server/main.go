package main

import (
	"context"
	"log"
	"nearest-charging-stations/internal/db"
	"nearest-charging-stations/internal/handler"
	"nearest-charging-stations/internal/repository"
	"net/http"
)

func main() {
	ctx := context.Background()
	connString := "postgres://postgres:postgre890@localhost:5432/postgres"

	pool, err := db.NewPool(ctx, connString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer pool.Close()

	repo := repository.NewChargingStationRepository(pool)
	stationHandler := handler.NewChargingStationHandler(repo)
	homeHandler := handler.NewHomeHandler()

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle(
		"GET /static/",
		http.StripPrefix(
			"/static/",
			fileServer,
		),
	)
	mux.HandleFunc("/", homeHandler.Home)
	mux.HandleFunc(
		"GET /api/stations/nearby",
		stationHandler.FindNearestStations,
	)

	log.Println("Server starting on :8082")

	if err := http.ListenAndServe(":8082", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}

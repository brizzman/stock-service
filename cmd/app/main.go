package main

import (
	"context"
	"fmt"
	"stock-service/internal/adapters/outbound/postgres"
	"stock-service/internal/infrastructure/config"
	"stock-service/internal/infrastructure/logger"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.GetConfig()
	
	log := logger.GetLogger()

	// Init repository store (with postgresql inside)
	store, err := postgres.NewStorage(ctx, *log, cfg.PG.URL)
	if err != nil {
		return fmt.Errorf("postgres.NewStorage failed: %w", err)
	}


	fmt.Println(store)
	// Init service manager

	// Init controllers
	return nil
}
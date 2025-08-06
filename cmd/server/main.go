package main

import (
	"fmt"

	"github.com/miscla/codebase-golang/internal/config"
	"github.com/miscla/codebase-golang/internal/database"
	"github.com/miscla/codebase-golang/internal/models"
	"github.com/miscla/codebase-golang/internal/router"
	"github.com/miscla/codebase-golang/pkg/logger"
)

func main() {
	cfg := config.Load()
	log := logger.New()

	db, err := database.ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	r := router.Setup(db, log, cfg)

	addr := fmt.Sprintf(":%s", cfg.AppPort)
	log.Infof("Server running at %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

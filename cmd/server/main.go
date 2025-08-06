package main

import (
	"fmt"
	"log"

	"codebase-golang/internal/router"
	"codebase-golang/pkg/config"
	"codebase-golang/pkg/database"
	"codebase-golang/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig()
	gin.SetMode(cfg.GinMode)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	logger.Init()
	defer logger.Sync()

	if err := database.Init(cfg); err != nil {
		logger.Log.Fatal("Database connection failed", zap.Error(err))
	}

	r := router.SetupRouter()
	addr := fmt.Sprintf(":%s", cfg.AppPort)

	logger.Log.Info("Starting server...",
		zap.String("app_name", cfg.AppName),
		zap.String("port", cfg.AppPort),
	)

	if err := r.Run(addr); err != nil {
		logger.Log.Fatal("Server failed", zap.Error(err))
	}
}

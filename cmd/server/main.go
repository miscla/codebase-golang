package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"codebase-golang/internal/router"
	"codebase-golang/internal/service"
	"codebase-golang/pkg/config"
	"codebase-golang/pkg/database"
	"codebase-golang/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	gin.SetMode(cfg.GinMode)

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("db connect error: %v", err)
	}
	defer db.Close()

	// create user service
	userSvc := service.NewUserService(db, cfg)

	// optionally pre-warm cache at startup
	if err := userSvc.RefreshCache(); err != nil {
		logger.Error("initial cache refresh failed:", err)
	}

	r := router.NewRouter(userSvc)

	// run server
	go func() {
		addr := ":" + cfg.AppPort
		logger.Info("server starting at", addr)
		if err := r.Run(addr); err != nil {
			logger.Error("server stopped:", err)
		}
	}()

	// graceful shutdown wait
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("shutting down server...")
	// allow a brief moment
	time.Sleep(1 * time.Second)
}

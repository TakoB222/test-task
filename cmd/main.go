package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"test-task/internal/config"
	"test-task/internal/delivery/http"
	"test-task/internal/server"
	"test-task/internal/service"
	"test-task/pkg/logger"
	"time"
)
import "github.com/sirupsen/logrus"

func init() {
	logrus.SetFormatter(new(logrus.TextFormatter))
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	cfg, err := initConfigManager()
	if err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	service := service.NewServices()
	handler := http.NewHandler(service)

	server := server.NewServer(server.Config{Host: cfg.Http.Host, Port: cfg.Http.Port,
		ReadTimeout: cfg.Http.ReadTimeout, WriteTimeout: cfg.Http.WriteTimeout}, handler.Init())
	go func() {
		if err := server.Run(); err != nil {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logrus.Info("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := server.Shutdown(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}
}

func initConfigManager() (*config.Config, error) {
	return config.Init("../config/config")
}

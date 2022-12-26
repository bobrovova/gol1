package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bobrovova/go-weather/pkg/service"
)

// @title Weather Api
// @version 1.0
// @description API Server for getting current weather

// @host localhost:8881
// @BasePath /

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	logger := service.NewLogger()

	service.StartServer(logger)
}

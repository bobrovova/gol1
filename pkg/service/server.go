package service

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/bobrovova/go-weather/configs"
	"github.com/bobrovova/go-weather/pkg/handler"
	"github.com/sirupsen/logrus"
)

func StartServer(log *logrus.Logger) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	router := handler.InitRoutes(log)

	log.Info("Starting application on port %d\n", configs.PORT)
	server := http.Server{
		Addr:         "localhost:" + strconv.Itoa(configs.PORT),
		WriteTimeout: 3 * time.Second,
		Handler:      router,
	}

	go func() {
		server.ListenAndServe()
	}()

	<-done
	log.Info("Server stopping...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		log.Writer().Close()
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Error(err)
	}

	log.Info("Server stopped")
}

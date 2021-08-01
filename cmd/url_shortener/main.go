package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"url_shortener/service"
	"url_shortener/service/api"
	"url_shortener/service/database"
	"url_shortener/service/domain"
	"url_shortener/service/logger"
)

func main() {
	var cfg service.Config
	if err := cfg.Load(); err != nil {
		panic("failed to load config: " + err.Error())
	}

	ctx := context.Background()

	log := logger.Setup(cfg.Log)

	ctx = logger.WithLogger(ctx, log)

	dbConn, err := database.Open(ctx, cfg.DB)
	if err != nil {
		log.WithError(err).Error("failed to open DB")
	}

	urlStorage := database.NewURLStorage(dbConn)
	shortURLs := domain.NewShortURLs(urlStorage, cfg.Domain.ShortURLsBase)
	handler := api.NewHandler(cfg.API, shortURLs)
	server := http.Server{
		Addr:    "0.0.0.0:" + cfg.Port,
		Handler: handler,
		BaseContext: func(net.Listener) context.Context {
			return logger.WithLogger(context.Background(), log)
		},
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		sig := make(chan os.Signal, 1)
		defer close(sig)
		signal.Notify(sig, os.Interrupt)
		<-sig
		if err := server.Shutdown(ctx); err != nil {
			log.WithError(err).Error("failed to shutdown HTTP server gracefully")
		} else {
			log.Info("HTTP server gracefully stopped")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Info("Starting HTTP server...")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.WithError(err).Error("failed to start HTTP server")
		}
	}()

	wg.Wait()

	if err := dbConn.Close(ctx); err != nil {
		log.WithError(err).Error("failed to close DB")
	}
}

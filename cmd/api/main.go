package main

import (
	"context"
	"errors"
	"media/internal/api"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/EmotivesProject/common/logger"
	"github.com/EmotivesProject/common/middlewares"
	"github.com/EmotivesProject/common/verification"
)

const timeBeforeTimeout = 15

func main() {
	logger.InitLogger("media", logger.EmailConfig{
		From:     os.Getenv("EMAIL_FROM"),
		Password: os.Getenv("EMAIL_PASSWORD"),
		Level:    os.Getenv("EMAIL_LEVEL"),
	})

	verification.Init(verification.VerificationConfig{
		VerificationURL: "http://uacl/authorize",
	})

	middlewares.Init(middlewares.Config{
		AllowedOrigins: os.Getenv("ALLOWED_ORIGINS"),
		AllowedMethods: "GET,POST,OPTIONS,DELETE",
		AllowedHeaders: "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token",
	})

	router := api.CreateRouter()

	srv := http.Server{
		Handler:      router,
		Addr:         os.Getenv("HOST") + ":" + os.Getenv("PORT"),
		WriteTimeout: timeBeforeTimeout * time.Second,
		ReadTimeout:  timeBeforeTimeout * time.Second,
	}

	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)

		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)

		<-sigint

		logger.Infof("Shutting down server")

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			logger.Infof("HTTP server Shutdown: %v", err)
		}

		close(idleConnsClosed)
	}()

	logger.Info("Starting Server")

	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		// Error starting or closing listener:
		logger.Infof("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}

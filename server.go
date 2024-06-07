package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const (
	_defaultReadTimeOut  = 15 * time.Second
	_defaultWriteTimeout = 30 * time.Second
	_defaultPort         = 8080
	_shutdownTimout      = 5 * time.Second
)

func RunHTTPServer(apiHandler func(chi.Router) http.Handler) {
	root := chi.NewMux()
	setMiddlewares(root)
	addCorsMiddleware(root)

	root.Mount("/api", apiHandler(root))

	port := _defaultPort
	if ps := os.Getenv("HTTP_PORT"); ps != "" {
		if prt, err := strconv.Atoi(ps); err == nil {
			port = prt
		}
	}

	srv := http.Server{
		Addr:         fmt.Sprintf(":%v", port),
		Handler:      root,
		ReadTimeout:  _defaultReadTimeOut,  // default values
		WriteTimeout: _defaultWriteTimeout, // default values
	}
	done := make(chan struct{})
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Starting HTTP server")
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			// errored out
			slog.Error("listenAndServe error - ", slog.String("err", err.Error()))
		}
		close(done)
	}()

	// setting up graceful shutdown
	select {
	case <-quit:
		// terminate called
		slog.Error("interrupt received, shutting down server")
	case <-done:
		// server errored out
	}

	ctx, shutdown := context.WithTimeout(context.Background(), _shutdownTimout)
	defer shutdown()

	err := srv.Shutdown(ctx)
	if err != nil {
		slog.Error("errored out while shutting down server")
	}
}

func setMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	addCorsMiddleware(router)

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)
}

func addCorsMiddleware(router *chi.Mux) {
	allowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ";")
	if len(allowedOrigins) == 0 {
		return
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(corsMiddleware.Handler)
}

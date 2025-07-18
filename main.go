package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"svc-subman/src/common/config"
	"svc-subman/src/ports_adapters/primary/http_server"
	"svc-subman/src/ports_adapters/secondary/persistence/db"
	"svc-subman/src/ports_adapters/secondary/service/application"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"time"
)

func main() {
	conf := config.InitEnv()
	logger := config.InitLogger(conf.LogLvl)

	repo, err := db.NewRepositories(conf.Repository)
	if err != nil {
		logger.Error(err.Error())
	}

	app, err := application.InitApp(repo.SubscribeRepo, logger)
	if err != nil {
		logger.Error(err.Error())
	}

	h := http_server.NewServer(app)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(120 * time.Second))
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Post("/subscribe", h.CreateSubscribe)
	r.Delete("/subscribe", h.DeleteSubscriber)
	r.Get("/subscribe/cost", h.AgreeSubscribe)
	r.Get("/subscribe/{id}", h.GetSubscribe)
	r.Get("/subscribes", h.GetSubscribes)
	server := &http.Server{Addr: ":8080", Handler: r}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		logger.Info("HTTP server run on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Error HTTP server: %v", err)
		}
	}()

	<-ctx.Done()
	logger.Info("Signal received to terminate, stop services...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("HTTP server termination error: %v", err)
	}
	logger.Info("HTTP server terminated")

}

package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
	"shortener/internal/config"
	urlHandlerPkg "shortener/internal/http/handler/url"
	"shortener/internal/repository/sqlite"
	urlRepositoryPkg "shortener/internal/repository/sqlite/url"
	urlServicePkg "shortener/internal/service/url"
	sl "shortener/internal/utils/logger/slog"
	"time"
)

func main() {
	// Load application configuration
	cfg := config.MustLoad()
	// Setup logger
	log := setupLogger(cfg.Env)

	log.Info("starting the application")

	// Create sqlite connection
	sqliteConnection, err := setupSqliteConnection(cfg.StoragePath)
	if err != nil {
		log.Error("failed to create sqlite connection", sl.Err(err))
		os.Exit(1)
	}

	// Close sqlite connection
	defer func(sqliteConnection *sql.DB) {
		err = closeSqliteConnection(sqliteConnection)
		if err != nil {
			log.Error("failed to close sqlite connection", sl.Err(err))
			os.Exit(1)
		}
	}(sqliteConnection)

	// Initialize URL repository/service/handler
	urlRepository := urlRepositoryPkg.NewRepository(sqliteConnection)
	urlService := urlServicePkg.NewService(urlRepository)
	urlHandler := urlHandlerPkg.NewHandler(urlService)

	// Setup router
	router := setupRouter(urlHandler)

	// Setup server
	server := &http.Server{
		Addr:           ":8000",
		Handler:        router,
		MaxHeaderBytes: 1 << 2,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	// Start server
	if err = server.ListenAndServe(); err != nil {
		log.Error("failed to start server", sl.Err(err))
	}
}

const (
	localEnv = "local"
	devEnv   = "dev"
	prodEnv  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case localEnv:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case devEnv:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case prodEnv:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}

func setupSqliteConnection(storagePath string) (*sql.DB, error) {
	return sqlite.NewConnection(storagePath)
}

func closeSqliteConnection(db *sql.DB) error {
	return sqlite.CloseConnection(db)
}

func setupRouter(urlHandler *urlHandlerPkg.Handler) *gin.Engine {
	ginEngine := gin.Default()

	api := ginEngine.Group("/api")
	{
		url := api.Group("/url")
		{
			url.GET("/:slug", urlHandler.GetURL)
			url.POST("/", urlHandler.SaveURL)
		}
	}

	return ginEngine
}

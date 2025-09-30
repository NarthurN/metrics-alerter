package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/NarthurN/metrics-alerter/internal/router"
	serviceMetrics "github.com/NarthurN/metrics-alerter/internal/service/metrics"
	"github.com/NarthurN/metrics-alerter/pkg/storage/metrics/storage"
)

// Приложение и его зависимости
type App struct {
	Storage *storage.MemStorage
	Server  *http.Server
	// Logger
}

func NewApp(addr string) *App {
	// инициализируем зависимости
	storage := storage.NewMemStorage()
	log.Println("✅ Инициализирован репозиторный слой metrics")

	service := serviceMetrics.NewMetricsService(storage)
	log.Println("✅ Инициализирован сервисный слой metrics")

	mux := router.NewRouter(service)
	log.Println("✅ Инициализирован роутер")

	srv := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	log.Println("✅ Инициализирован server")

	log.Println("✅Инициализация приложения прошла успешно")
	return &App{
		Storage: storage,
		Server:  srv,
	}
}

func (a *App) Run(addr string) error {
	log.Println("🚀 Сервер слушает на порту", addr)
	if err := a.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (a *App) Shutdown(ctx context.Context) error {
	if err := a.Server.Shutdown(ctx); err != nil {
		return err
	}
	a.close()
	return nil
}

// Закрытие ресурсов
func (a *App) close() {
	log.Println("✅ Закрыли ресурсы")
}

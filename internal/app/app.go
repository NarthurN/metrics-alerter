package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NarthurN/metrics-alerter/internal/repository"
	repoMetrics "github.com/NarthurN/metrics-alerter/internal/repository/metrics"
	"github.com/NarthurN/metrics-alerter/internal/router"
	serviceMetrics "github.com/NarthurN/metrics-alerter/internal/service/metrics"
)

// Приложение и его зависимости
type App struct {
	Storage repository.ServerRepository
	Server  *http.Server
	// Logger
}

func NewApp() *App {
	// инициализируем зависимости
	storage := repoMetrics.NewMemStorage()
	log.Println("✅ Инициализирован репозиторный слой metrics")

	service := serviceMetrics.NewMetricsService(storage)
	log.Println("✅ Инициализирован сервисный слой metrics")

	mux := router.NewRouter(service)
	log.Println("✅ Инициализирован роутер")

	srv := &http.Server{
		Addr:         `:8080`,
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

func (a *App) Run() {
	log.Println("🚀 Сервер слушает на порту", a.Server.Addr)

	go func() {
		if err := a.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen and serve error: %v", err)
		}
	}()

	// Ожидаем сигнал для начала graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("✅ Получили сигнал о завершении работы")

	// Создаем контекст с таймаутом для shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.Server.Shutdown(ctx); err != nil {
		log.Printf("🏮 graceful shutdown error: %v", err)
	}
	a.Close()
}

// Закрытие ресурсов
func (a *App) Close() {
	log.Println("✅ Закрыли ресурсы")
}

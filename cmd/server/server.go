package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NarthurN/metrics-alerter/internal/app"
)

func main() {
	// Парсируем флаги
	parseFlags()

	// Инициализация зависимостей приложения
	application := app.NewApp(flagRunAddr)

	// Запуск приложения
	go func() {
		if err := application.Run(); err != nil {
			log.Fatalf("Ошибка при запуске http-сервера: %v", err)
		}
	}()

	// Ожидаем сигнал для начала graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("✅ Получили сигнал о завершении работы, начинаем graceful shutdown")

	// Создаем контекст с таймаутом для shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := application.Shutdown(ctx); err != nil {
		log.Fatalf("🏮 Ошибка graceful shutdown: %v", err)
	}
	log.Println("✅ Сервер успешно остановлен")
}

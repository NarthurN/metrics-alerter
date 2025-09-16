package main

import (
	"github.com/NarthurN/metrics-alerter/internal/app"
)

func main() {
	// Инициализация зависимостей приложения
	app := app.NewApp()

	// Запуск приложения
	app.Run()
}

package app

import (
	"log"

	"github.com/NarthurN/metrics-alerter/internal/repository/server"
)

// Приложение и его зависимости
type App struct {
	Storage server.ServerRepository
	// Logger
}

func NewApp() *App {
	// инициализируем зависимости
	log.Println("Инициализация приложения прошла успешно")
	return &App{}
}

// Закрытие ресурсов
func (a *App) Close() {
	log.Println("Закрыли ресурсы")
}

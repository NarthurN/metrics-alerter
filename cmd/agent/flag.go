package main

import (
	"flag"
)

var flagRunAddr string
var pollInterval int
var reportInterval int

func parseFlags() {
	flag.StringVar(&flagRunAddr, "a", "localhost:8080", "address and port to run server")
	flag.IntVar(&reportInterval, "r", 10, "частота отправки метрик на сервер в секундах")
	flag.IntVar(&pollInterval, "p", 2, "частота опроса метрик из пакета runtime в секундах")

	flag.Parse()
}

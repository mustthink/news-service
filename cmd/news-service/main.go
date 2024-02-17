package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mustthink/news-service/internal/service"
)

func main() {
	app := service.New()
	go app.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	app.Stop()
}

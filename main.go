package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/muhimron90/orders-api/application"
	"github.com/muhimron90/orders-api/config"
	l "github.com/muhimron90/orders-api/logging"
)

func main() {
	app := application.New(config.LoadConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	err := app.Start(ctx)
	if err != nil {
		l.Debug(true)
		l.Logger(err.Error())
	}

}

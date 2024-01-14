package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/muhimron90/orders-api/application"
	l "github.com/muhimron90/orders-api/logging"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	err := app.Start(ctx)
	if err != nil {
		l.Debug(true)
		l.Logger(err.Error())
	}

}

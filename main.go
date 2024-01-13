package main

import (
	"context"
	"github.com/muhimron90/orders-api/application"
	l "github.com/muhimron90/orders-api/logging"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		l.Debug(true)
		l.Logger(err.Error())
	}

}

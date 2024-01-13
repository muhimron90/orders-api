package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/muhimron90/orders-api/config"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: loadRoutes(),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    config.PORT_NUMBER,
		Handler: a.router,
	}
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server : %w", err)
	}
	return nil
}

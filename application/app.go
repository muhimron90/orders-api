package application

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/muhimron90/orders-api/config"
	"github.com/muhimron90/orders-api/logging"
	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
	config config.Config
}

func New(config config.Config) *App {
	app := &App{
		rdb: redis.NewClient(&redis.Options{
			Addr: config.RedisAddress,
		}),
		config: config,
	}
	app.loadRoutes()
	return app
}

func (a *App) Start(ctx context.Context) error {
	logging.Debug(true)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.ServerPort),
		Handler: a.router,
	}
	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect redis server : %w", err)
	}
	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()
	logging.Logger("starting server")

	var sPort string = strconv.FormatUint(uint64(a.config.ServerPort), 10)
	logging.Debug(true)
	logging.Logger(a.config.RedisAddress)
	logging.Logger(string(sPort))
	//create channel to concuren server
	ch := make(chan error, 1)
	//concurency server and assign error into channel
	go func() {
		//serve and listen server
		// err is callback from server
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server : %w", err)
		}
		// if error wasnt found close channel
		close(ch)
	}()
	// switch
	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}

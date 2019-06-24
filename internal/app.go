package internal

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/kevburnsjr/stupid-poker/internal/config"
	"github.com/kevburnsjr/stupid-poker/internal/controller"
	"github.com/kevburnsjr/stupid-poker/internal/poker"
)

type App struct {
	config    *config.App
	logger    *logrus.Logger
	server    *http.Server
	gameCache poker.GameCache
}

func NewApp(cfg *config.App) *App {
	return &App{
		config:    cfg,
		logger:    newLogger(cfg.Log.Level),
		gameCache: poker.NewGameCache(),
	}
}

func (app *App) Start() {
	cfg := app.config.Api

	handler := controller.NewRouter(app.logger, app.gameCache)

	app.server = &http.Server{
		Handler: handler,
		Addr:    ":" + cfg.Port,
	}
	go func() {
		var err error
		app.logger.Printf("App Listening on port %s", cfg.Port)
		if cfg.Ssl.Enabled {
			err = app.server.ListenAndServeTLS(cfg.Ssl.Cert, cfg.Ssl.Key)
		} else {
			err = app.server.ListenAndServe()
		}
		if err != nil {
			app.logger.Fatal(err.Error())
		}
	}()
}

func (app *App) Stop(timeout time.Duration) {
	app.logger.Printf("Stopping HTTP Listener")
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	app.server.Shutdown(ctx)
}

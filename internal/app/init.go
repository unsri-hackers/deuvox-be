package app

import (
	"context"
	"database/sql"
	"deuvox/pkg/handler"
	"deuvox/pkg/response"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

type AppConfig struct {
	Host            string
	Port            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

type App struct {
	Config     AppConfig
	delivery   delivery
	usecase    usecase
	repository repository
}

func New(serverCfg AppConfig, db *sql.DB) App {

	var app App
	app.initRepository(db)
	app.initUsecase()
	app.initDelivery()
	app.Config = serverCfg
	return app
}

func (app *App) createHandlers() http.Handler {
	r := chi.NewRouter()
	d := app.delivery

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.NotFound(handler.NotFound)
	r.MethodNotAllowed(handler.MethodNotAllowed)

	r.Group(func(r chi.Router) {
		// using jwt verifier
	})

	r.Group(func(r chi.Router) {
		r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
			response.Write(rw, http.StatusOK, "Hi", nil, "")
		})

		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", d.auth.Register)
			r.Post("/login", d.auth.Login)
		})
	})

	return r
}

func (app *App) StartServer() {
	osSignalChan := make(chan os.Signal, 1)
	signal.Notify(osSignalChan, os.Interrupt, syscall.SIGTERM)
	defer func() {
		signal.Stop(osSignalChan)
		os.Exit(0)
	}()

	r := app.createHandlers()
	address := fmt.Sprintf("%s:%s", app.Config.Host, app.Config.Port)
	server := &http.Server{
		Addr:         address,
		ReadTimeout:  app.Config.ReadTimeout,
		WriteTimeout: app.Config.WriteTimeout,
		Handler:      r,
	}

	shutdownCtx := context.Background()
	if app.Config.ShutdownTimeout > 0 {
		var cancelShutdownTimeout context.CancelFunc
		shutdownCtx, cancelShutdownTimeout = context.WithTimeout(shutdownCtx, app.Config.ShutdownTimeout)
		defer cancelShutdownTimeout()
	}
	log.Info().Msg(fmt.Sprintf("serving %s\n", address))
	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatal().Err(err).Stack().Msg("cannot start server")
	}

	go func(srv *http.Server) {
		<-osSignalChan
		err := srv.Shutdown(shutdownCtx)
		if err != nil {
			panic("failed to shutdown gracefully")
		}
	}(server)
}

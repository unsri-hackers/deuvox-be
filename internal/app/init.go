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

	"deuvox/pkg/log"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
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
	log.Info().Msg("Initilize repository")
	app.initRepository(db)
	log.Info().Msg("Initilize usecase")
	app.initUsecase()
	log.Info().Msg("Initilize delivery")
	app.initDelivery()
	app.Config = serverCfg
	return app
}

func (app *App) createHandlers() http.Handler {
	log.Info().Msg("Initilize router and handler")

	r := chi.NewRouter()
	d := app.delivery

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

		r.Get("/token", d.auth.Token)
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
	log.Info().Msgf("Listening server on %s", address)
	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Panic().Err(err).Msg("Failed to start server")
	}

	go func(srv *http.Server) {
		<-osSignalChan
		err := srv.Shutdown(shutdownCtx)
		if err != nil {
			log.Panic().Err(err).Msg("failed to shutdown gracefully")
		}
	}(server)
}

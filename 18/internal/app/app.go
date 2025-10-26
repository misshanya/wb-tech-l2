package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/misshanya/wb-tech-l2/18/internal/config"

	repo "github.com/misshanya/wb-tech-l2/18/internal/repository/event"
	service "github.com/misshanya/wb-tech-l2/18/internal/service/event"
	handler "github.com/misshanya/wb-tech-l2/18/internal/transport/http/v1/event"
)

type App struct {
	e   *echo.Echo
	cfg *config.Config
	l   *slog.Logger
}

// New creates and initializes a new instance of App
func New(cfg *config.Config, l *slog.Logger) (*App, error) {
	a := &App{
		cfg: cfg,
		l:   l,
	}

	repo := repo.New()
	s := service.New(
		l,
		repo,
	)
	h := handler.New(s)

	a.initEcho()

	apiGroup := a.e.Group("/api/v1")
	eventGroup := apiGroup.Group("/event")

	h.Setup(eventGroup)

	return a, nil
}

// Start performs a start of all functional services
func (a *App) Start(errChan chan<- error) {
	a.l.Info("Starting...",
		"addr", a.cfg.Server.Addr,
	)
	if err := a.e.Start(a.cfg.Server.Addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
		errChan <- err
	}
}

func (a *App) Stop(ctx context.Context) error {
	a.l.Info("[!] Shutting down...")

	var stopErr error

	a.l.Info("Stopping http server...")
	if err := a.e.Shutdown(ctx); err != nil {
		stopErr = errors.Join(stopErr, fmt.Errorf("failed to stop http server: %w", err))
	}

	if stopErr != nil {
		return stopErr
	}

	a.l.Info("Stopped gracefully!")
	return nil
}

// initEcho sets up a new Echo instance with logger
func (a *App) initEcho() {
	a.e = echo.New()
	a.e.HideBanner = true
	a.e.HidePort = true
	a.e.Pre(middleware.RemoveTrailingSlash())

	a.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{a.cfg.Server.CORSOrigin},
	}))

	a.e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				a.l.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("ip", v.RemoteIP),
					slog.String("latency", time.Now().Sub(v.StartTime).String()),
				)
			} else {
				a.l.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("ip", v.RemoteIP),
					slog.String("latency", time.Now().Sub(v.StartTime).String()),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))

	a.e.Use(middleware.Recover())
}

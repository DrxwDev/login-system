package http

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	"github.com/DrxwDev/login-system/internal/platform/config"
)

func NewServer(lc fx.Lifecycle, app config.AppConfig, srv config.ServerConfig, router *gin.Engine, logger *slog.Logger) *http.Server {
	server := &http.Server{
		Addr:              net.JoinHostPort(app.HOST, app.PORT),
		Handler:           router,
		ReadTimeout:       srv.ReadTimeout,
		ReadHeaderTimeout: srv.ReadHeaderTimeout,
		WriteTimeout:      srv.WriteTimeout,
		IdleTimeout:       srv.IdleTimeout,
		MaxHeaderBytes:    srv.MaxHeaderBytes,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Server runnig at", "addr", server.Addr)
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Error("Failed to start server", "err", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Shuting down server")
			context, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			if err := server.Shutdown(context); err != nil {
				logger.Error("Server Shuting down error", "err", err)
				return err
			}

			logger.Info("Server Shutdown")
			return nil
		},
	})

	return server
}

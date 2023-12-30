package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-thread-7/lacoste-romberg-d1807/app/api-gateway/http-server"
	"github.com/go-thread-7/lacoste-romberg-d1807/app/app/middlewares/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func RunServers(lc fx.Lifecycle, ctx context.Context, cfg *config.Config, log logger.ILogger, e *echo.Echo, grpcServer *grpc.GrpcServer) error {

	lc.Append(
		fx.Hook{
			OnStart: func(_ context.Context) error {
				go func() {
					if err := http_server.RunHttpServer(ctx, e, log, cfg.Echo); !errors.Is(err, http.ErrServerClosed) {
						log.Fatalf("error running http server: %v", err)
					}
				}()

				go func() {
					if err := grpcServer.RunGrpcServer(ctx); !errors.Is(err, http.ErrServerClosed) {
						log.Fatalf("error running grpc server: %v", err)
					}
				}()

				e.GET("/", func(c echo.Context) error {
					return c.String(http.StatusOK, config.GetMicroserviceName(cfg.ServiceName))
				})

				return nil
			},
			OnStop: func(_ context.Context) error {
				log.Infof("all servers shutdown gracefully...")
				return nil
			},
		},
	)

	return nil
}

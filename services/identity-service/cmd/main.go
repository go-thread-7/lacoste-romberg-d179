package main

import (
	"github.com/go-thread-7/lacoste-romberg-d1807/app/middlewares/logger"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				logger.InitLogger,
			),
		),
	).Run()
}

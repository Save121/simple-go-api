package main

import (
	"context"
	"fmt"

	"github.com/Save121/simple-go-api/database"
	"github.com/Save121/simple-go-api/internal/api"
	"github.com/Save121/simple-go-api/internal/repository"
	"github.com/Save121/simple-go-api/internal/service"
	"github.com/Save121/simple-go-api/settings"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
			api.New,
			echo.New,
		),
		fx.Invoke(
			setLifeCycle,
		),
	)
	app.Run()
}
func setLifeCycle(lc fx.Lifecycle, api *api.API, s *settings.Settings, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%d", s.Port)
			go api.Start(e, address)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}

package main

import (
	"context"

	"github.com/Save121/simple-go-api/database"
	"github.com/Save121/simple-go-api/internal/repository"
	"github.com/Save121/simple-go-api/internal/service"
	"github.com/Save121/simple-go-api/settings"
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
		),
		fx.Invoke(),
	)
	app.Run()
}

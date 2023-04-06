package main

import (
	"github.com/Save121/simple-go-api/settings"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			settings.New,
		),
		fx.Invoke(),
	)
	app.Run()
}

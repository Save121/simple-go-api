package api

import (
	"github.com/Save121/simple-go-api/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type API struct {
	serv          service.Service
	dataValidator *validator.Validate
}

func New(serv service.Service) *API {
	return &API{
		serv:          serv,
		dataValidator: validator.New(),
	}
}

func (api *API) Start(e *echo.Echo, address string) error {
	api.RegisterRoutes(e)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.POST, echo.PUT, echo.GET, echo.DELETE},
		AllowHeaders:     []string{echo.HeaderContentType},
		AllowCredentials: true,
	}))
	return e.Start(address)
}

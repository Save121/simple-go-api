package api

import (
	"github.com/Save121/simple-go-api/internal/service"
	"github.com/labstack/echo/v4"
)

type API struct {
	serv service.Service
}

func New(serv service.Service) *API {
	return &API{
		serv: serv,
	}
}

func (api *API) Start(e *echo.Echo, address string) error {
	api.RegisterRoutes(e)
	return e.Start(address)
}

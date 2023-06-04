package api

import (
	"errors"
	"net/http"

	"github.com/Save121/simple-go-api/internal/api/dtos"
	"github.com/Save121/simple-go-api/internal/service"
	"github.com/labstack/echo/v4"
)

func (api *API) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterUser{}
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = api.serv.RegisterUser(ctx, params.Email, params.Name, params.Password)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, err)
		}
		return c.JSON(http.StatusInternalServerError, errors.New("Internal server error"))
	}
	return c.JSON(http.StatusCreated, nil)
}

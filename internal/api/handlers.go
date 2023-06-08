package api

import (
	"net/http"

	"github.com/Save121/simple-go-api/internal/api/dtos"
	"github.com/Save121/simple-go-api/internal/models"
	"github.com/Save121/simple-go-api/internal/service"
	"github.com/Save121/simple-go-api/jwt"
	"github.com/labstack/echo/v4"
)

type responseMessage struct {
	Message string `json:"message"`
}

func (api *API) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterUser{}
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}
	err = api.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})

	}

	err = api.serv.RegisterUser(ctx, params.Email, params.Name, params.Password)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "User already exists"})
		}
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}
	return c.JSON(http.StatusCreated, nil)
}

func (api *API) loginUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.LoginUser{}
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}
	err = api.dataValidator.Struct(params)

	u, err := api.serv.LoginUser(ctx, params.Email, params.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})

	}
	token, err := jwt.SignedLoginToken(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}
	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Path:     "/",
	}

	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, map[string]string{"success": "true"})

}

func (api *API) AddMovie(c echo.Context) error {
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "Unauthorized"})
	}

	claims, err := jwt.ParseLoginJWT(cookie.Value)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "Unauthorized"})
	}

	email := claims["email"].(string)
	params := dtos.AddMovie{}
	ctx := c.Request().Context()

	err = c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	err = api.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}
	m := models.Movie{
		Name:        params.Name,
		Description: params.Description,
		Price:       params.Price,
	}
	err = api.serv.AddMovie(ctx, m, email)
	if err != nil {

		if err == service.ErrInvalidPermissions {
			return c.JSON(http.StatusForbidden, responseMessage{Message: "Invalid permissions"})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	return c.JSON(http.StatusCreated, nil)
}

package controllers

import (
	"echo/auth/entity"
	"echo/auth/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	S service.AuthService
}

func (c *Controller) Register(ctx echo.Context) error {
	// entity user
	user := entity.User{}
	// request body
	err := ctx.Bind(&user)
	// if request body is empty, return error
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "request body is empty",
			"error":   err.Error(),
		})
	}
	error := ctx.Validate(&user)
	if error != nil {
		return error
	}
	// if request body is not empty, call service.Register
	er := c.S.Register(user)
	// if service.Register error, return error
	if er != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "error",
			"error":   er.Error(),
		})
	}
	// if service.Register success, return success
	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "success",
		"data":    user,
	})
}

func (c *Controller) Login(ctx echo.Context) error {
	// entity user
	user := entity.User{}
	// request body
	err := ctx.Bind(&user)
	// if request body is empty, return error
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "request body is empty",
			"error":   err.Error(),
		})
	}

	// if request body is not empty, call service.Login
	users, er := c.S.Login(user.Email, user.Password)
	// if service.Login error, return error
	if er != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "error",
			"error":   er.Error(),
		})
	}
	// if service.Login success, return success
	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "success",
		"data":    users,
	})
}

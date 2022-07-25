package main

import (
	"echo/auth/config"
	"echo/auth/helper"
	"echo/auth/routers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Validator = &helper.CustomValidator{Validator: validator.New()}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	conf := config.DBConfig{}

	routers.AuthRoutes(e, conf)

	err := e.Start("127.0.0.1:4000")
	if err != nil {
		panic(err)
	}
}

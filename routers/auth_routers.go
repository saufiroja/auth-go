package routers

import (
	"echo/auth/config"
	"echo/auth/controllers"
	"echo/auth/repository"
	"echo/auth/service"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(echo *echo.Echo, conf config.DBConfig) {
	db := config.InitDB(conf)

	repo := repository.NewAuthRepository(db)
	service := service.NewAuthService(repo, conf)
	control := controllers.Controller{
		S: service,
	}

	echo.POST("/api/register", control.Register)
	echo.POST("/api/login", control.Login)
}

package account

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lilihx/chatRoom/common/config"
	"github.com/lilihx/chatRoom/common/echotool"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitServerAndStart() {
	e := echo.New()
	echotool.RegisterPProf(e)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("index", index)
	e.GET("oauth/redirect", oauth)

	e.Logger.Fatal(e.Start(config.Config.AccountServer.Addr))
}

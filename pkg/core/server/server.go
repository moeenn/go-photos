package server

import (
	"github.com/labstack/echo/v4"
	"webapp/pkg/core/config"
	"webapp/pkg/core/config/templates"
)

func New(config *config.Config) *echo.Echo {
	server := echo.New()
	server.Static(config.StaticPrefix, config.StaticPath)
	server.Renderer = templates.New(config.ViewsPath)

	return server
}

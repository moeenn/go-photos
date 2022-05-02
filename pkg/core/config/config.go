package config

import (
	"webapp/pkg/core/config/env"
)

type Config struct {
	ServerPort   string
	ViewsPath    string
	StaticPath   string
	StaticPrefix string
}

func New() (*Config, error) {
	port, err := env.Get("SERVER_PORT")
	if err != nil {
		return &Config{}, err
	}

	return &Config{
		ServerPort:   port,
		ViewsPath:    "views/*.html",
		StaticPath:   "assets/",
		StaticPrefix: "/static/",
	}, nil
}

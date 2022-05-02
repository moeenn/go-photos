package main

import (
  "fmt"
  "webapp/pkg/core/config"
  "webapp/pkg/core/server"
  "webapp/pkg/routes"
)

func main() {
  config, err := config.New()
  if err != nil {
    fmt.Printf("[config error] %s\n", err)
    return
  }

  srv := server.New(config)
  routes.Register(srv)

  srv.Logger.Fatal(srv.Start(":" + config.ServerPort))
}

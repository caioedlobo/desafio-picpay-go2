package main

import (
	"desafio-picpay-go2/internal/config"
	"desafio-picpay-go2/internal/infra/container"
	"desafio-picpay-go2/internal/infra/http/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	env := config.GetConfig()
	e := echo.New()
	c := container.NewContainer(env)
	handler.RegisterHandler(e, c)
	err := e.Start(env.Port)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"desafio-picpay-go2/internal/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	handler.RegisterHandler(e)
	err := e.Start(":1323")
	if err != nil {
		panic(err)
	}
}

package main

import (
	"context"
	"database/sql"
	"desafio-picpay-go2/internal/domain/user"
	"desafio-picpay-go2/internal/domain/user/usecase"
	"desafio-picpay-go2/internal/infra/container"
	"desafio-picpay-go2/internal/infra/http/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	ctx := context.Background()
	c := container.NewContainer(ctx)
	handler.RegisterHandler(e, c)
	err := e.Start(":1323")
	if err != nil {
		panic(err)
	}
}

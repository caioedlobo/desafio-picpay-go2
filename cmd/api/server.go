package main

import (
	"desafio-picpay-go2/internal/config"
	"desafio-picpay-go2/internal/infra/container"
	"desafio-picpay-go2/internal/infra/http/handler"
	logger2 "desafio-picpay-go2/internal/infra/logger"
	"github.com/labstack/echo/v4"
	"os"
	"runtime/debug"
)

func main() {
	env := config.GetConfig()
	logger := logger2.NewLogger(env)
	defer func() {
		if err := recover(); err != nil {
			logger.Error("panic",
				"panic", err,
				"stack", string(debug.Stack()),
			)
		}
	}()
	cont, err := container.NewContainer(env)
	if err != nil {
		logger.Fatal("failed to initialize container", "error", err)
		os.Exit(1)
	}
	ec := echo.New()

	handler.RegisterHandler(ec, cont)
	err = ec.Start(env.Port)
	if err != nil {
		panic(err)
	}
}

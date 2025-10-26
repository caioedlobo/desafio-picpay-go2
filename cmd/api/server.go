package main

import (
	"desafio-picpay-go2/internal/config"
	"desafio-picpay-go2/internal/infra/container"
	"desafio-picpay-go2/internal/infra/http/handler"
	"desafio-picpay-go2/internal/infra/http/middleware"
	logger2 "desafio-picpay-go2/internal/infra/logger"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	cont, err := container.NewContainer(env, logger)
	if err != nil {
		logger.Fatal("failed to initialize container", "error", err)
		os.Exit(1)
	}
	ec := echo.New()
	middleware.Apply(ec, middleware.Config{Metrics: cont.Metrics})
	ec.GET("/metrics", echo.WrapHandler(promhttp.HandlerFor(cont.Metrics.Registry(), promhttp.HandlerOpts{})))
	handler.RegisterHandler(ec, cont)
	err = ec.Start(env.Port)
	if err != nil {
		panic(err)
	}
}

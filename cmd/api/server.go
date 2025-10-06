package main

import (
	"database/sql"
	"desafio-picpay-go2/internal/domain/user"
	"desafio-picpay-go2/internal/domain/user/usecase"
	"desafio-picpay-go2/internal/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	var db *sql.DB
	userRepo := user.NewRepository(db)
	usecase.NewUserCase(userRepo)
	handler.RegisterHandler(e)
	err := e.Start(":1323")
	if err != nil {
		panic(err)
	}
}

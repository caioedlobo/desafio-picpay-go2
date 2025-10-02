package main

import (
	"desafio-picpay-go2/internal/handler"
)

func main() {
	e := handler.NewHandler()
	err := e.Listen("1323")
	if err != nil {
		panic(err)
	}
}

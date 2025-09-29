package main

import (
	"desafio-picpay-go2/internal/handlers"
)

func main() {
	e := handlers.NewHandler()
	err := e.Listen("1323")
	if err != nil {
		panic(err)
	}
}

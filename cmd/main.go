package main

import (
	"desafio-picpay-go2/internal/domain/user"
	"desafio-picpay-go2/internal/domain/user/value_object"
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	password, err := value_object.NewPassword("pass123")
	if err != nil {
		panic(err)
	}
	user, err := user.NewUser("Caio", "96334365584", "cpf2", "caioeduardolobo@gmail.com", *password)
	if err != nil {
		panic(err)
	}
	fmt.Println(user.DocumentType)
}

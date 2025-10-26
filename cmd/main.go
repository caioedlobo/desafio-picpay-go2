package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

/*func main() {
	data := "exa"
	err := validation.Validate(data,
		validation.Required,       // not empty
		validation.Length(5, 100), // length between 5 and 100
		is.URL,                    // is a valid URL
	)
	fmt.Println(err)
	// Output:
	// must be a valid URL
}*/

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,lte=100"`
}

func main() {
	validate := validator.New()

	req := LoginRequest{
		Email:    "email-invalido",
		Password: "",
	}

	err := validate.Struct(req)
	if err != nil {
		var errorList []string

		for _, e := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("Campo '%s' falhou na regra '%s'", e.Field(), e.Tag())
			errorList = append(errorList, msg)
		}

		fmt.Println("Erros de validação:")
		for _, msg := range errorList {
			fmt.Println("-", msg)
		}
	}
}

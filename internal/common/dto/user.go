package dto

import "github.com/shopspring/decimal"

type CreateUserRequest struct {
	Name           string `json:"name" validate:"required,lte=100"`
	DocumentNumber string `json:"document_number" validate:"required,gte=11,lte=14"`
	DocumentType   string `json:"document_type" validate:"required,gte=3,lte=4"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,lte=100"`
}

type CreateUserResponse struct {
	Id             string `json:"id" validate:"required"`
	Name           string `json:"name" validate:"required,lte=100"`
	DocumentNumber string `json:"document_number" validate:"required,gte=11,lte=14"`
	DocumentType   string `json:"document_type" validate:"required,gte=3,lte=4"`
	Email          string `json:"email" validate:"required,email"`
}

type UserResponse struct {
	Name           string          `json:"name" validate:"required,lte=100"`
	DocumentNumber string          `json:"document_number" validate:"required,gte=11,lte=14"`
	Balance        decimal.Decimal `json:"balance" validate:"required"`
	Email          string          `json:"email" validate:"required,email"`
}

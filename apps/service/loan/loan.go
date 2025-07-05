package loan

import (
	"multifinancetest/apps/repositories"

	"github.com/go-playground/validator/v10"
)

type loan struct {
	db        repositories.IDatabase
	validator *validator.Validate
}

func NewLoan(
	db repositories.IDatabase,
	validator *validator.Validate,
) *loan {
	return &loan{
		db:        db,
		validator: validator,
	}
}

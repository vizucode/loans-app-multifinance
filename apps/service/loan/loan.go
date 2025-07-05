package loan

import (
	"multifinancetest/apps/repositories"
)

type Loan struct {
	db repositories.IDatabase
}

func NewLoan(db repositories.IDatabase) *Loan {
	return &Loan{
		db: db,
	}
}

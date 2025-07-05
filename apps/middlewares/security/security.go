package security

import "multifinancetest/apps/repositories"

type security struct {
	db repositories.IDatabase
}

func NewSecurity(
	db repositories.IDatabase,
) *security {
	return &security{
		db: db,
	}
}

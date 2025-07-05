package rest

import (
	"multifinancetest/apps/domain"
	"multifinancetest/apps/middlewares/security"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vizucode/gokit/logger"
)

func (r *rest) CreateLoan(c *fiber.Ctx) error {

	ctx, _, _ := security.ExtractUserContextFiber(c)

	var payload domain.RequestLoans
	err := c.BodyParser(&payload)
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	resultLoan, err := r.loanService.CreateLoan(ctx, payload)
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return r.ResponseJson(c,
		http.StatusOK,
		resultLoan,
		"Successfully Create Loan",
		domain.Metadata{},
	)
}

func (r *rest) GetCustomerLimitLoan(c *fiber.Ctx) error {

	ctx, _, _ := security.ExtractUserContextFiber(c)

	resultCustomerLoan, err := r.loanService.GetLimitLoans(ctx)
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return r.ResponseJson(c,
		http.StatusOK,
		resultCustomerLoan,
		"Successfully Get Loan",
		domain.Metadata{},
	)
}

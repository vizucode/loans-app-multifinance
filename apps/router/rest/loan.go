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

func (r *rest) GetCustomerHistoryLoan(c *fiber.Ctx) error {

	ctx, _, _ := security.ExtractUserContextFiber(c)

	filter := domain.Filter{
		Limit: c.QueryInt("limit", 10),
		Page:  c.QueryInt("page"),
		Where: map[string]any{
			"search": c.Query("search"),
		},
	}

	resultCustomerLoan, err := r.loanService.GetHistoryLoans(ctx, filter)
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	countHistoryLoans, err := r.loanService.GetHistoryLoans(ctx, domain.Filter{
		Where: map[string]any{
			"search": c.Query("search"),
		},
	})
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	var totalPage int
	if filter.Limit == 0 {
		totalPage = 1
	} else {
		totalPage = len(countHistoryLoans) / filter.Limit
	}

	return r.ResponseJson(c,
		http.StatusOK,
		resultCustomerLoan,
		"Successfully Get Loan",
		domain.Metadata{
			TotalItem:     int64(len(countHistoryLoans)),
			TotalPage:     totalPage,
			TotalPageItem: len(resultCustomerLoan),
		},
	)
}

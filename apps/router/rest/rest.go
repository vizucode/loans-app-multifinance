package rest

import (
	"strconv"

	"multifinancetest/apps/domain"
	"multifinancetest/apps/middlewares"
	"multifinancetest/apps/service"

	"github.com/gofiber/fiber/v2"
)

type rest struct {
	mw          middlewares.IMiddleware
	authService service.IAuthService
	loanService service.ILoans
}

func NewRest(
	mw middlewares.IMiddleware,
	authService service.IAuthService,
	loanService service.ILoans,
) *rest {
	return &rest{
		mw:          mw,
		authService: authService,
		loanService: loanService,
	}
}

func (r *rest) Router(app fiber.Router) {
	v1 := app.Group("/v1")

	// auth endpoint
	v1.Post("/signup", r.SignUp)
	v1.Post("/signin", r.SignIn)
	v1.Post("/refresh-token", r.RefreshToken)
	v1.Get("/customer/profile", r.mw.AuthMiddleware, r.GetProfile)
	v1.Post("/signout", r.mw.AuthMiddleware, r.Signout)

	// loan endpoint
	v1.Post("/customer/loan", r.mw.AuthMiddleware, r.CreateLoan)
}

func (rest *rest) ResponseJson(
	ctx *fiber.Ctx,
	StatusCode int,
	data interface{},
	message string,
	metadata domain.Metadata,
) error {
	return ctx.Status(StatusCode).JSON(&domain.ResponseJson{
		MetaData:   metadata,
		StatusCode: strconv.Itoa(StatusCode),
		Data:       data,
		Message:    message,
	})
}

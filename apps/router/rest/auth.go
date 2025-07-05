package rest

import (
	"net/http"

	"multifinancetest/apps/domain"
	"multifinancetest/apps/middlewares/security"
	"multifinancetest/helpers/constants/httpstd"
	"multifinancetest/helpers/constants/rpcstd"

	"github.com/gofiber/fiber/v2"
	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
)

func (r *rest) SignIn(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var payload domain.RequestSignIn
	err := c.BodyParser(&payload)
	if err != nil {
		logger.Log.Error(ctx, err)
		return errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.ABORTED, httpstd.BadRequest)
	}

	resultSignIn, err := r.authService.SignIn(ctx, payload)
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return r.ResponseJson(c,
		http.StatusOK,
		resultSignIn,
		"Successfully Sign In",
		domain.Metadata{},
	)
}

func (r *rest) RefreshToken(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var payload domain.RequestRefreshToken
	err := c.BodyParser(&payload)
	if err != nil {
		logger.Log.Error(ctx, err)
		return errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.ABORTED, httpstd.BadRequest)
	}

	resultRefreshToken, err := r.authService.RefreshToken(ctx, payload.AccessToken, payload.RefreshToken)
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return r.ResponseJson(c,
		http.StatusOK,
		resultRefreshToken,
		"Successfully Refresh Token",
		domain.Metadata{},
	)
}

func (r *rest) Signout(c *fiber.Ctx) error {
	ctx, _, _ := security.ExtractUserContextFiber(c)

	err := r.authService.SignOut(ctx)
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return r.ResponseJson(c,
		http.StatusOK,
		nil,
		"Successfully Sign Out",
		domain.Metadata{},
	)
}

func (r *rest) SignUp(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var payload domain.RequestCustomer
	err := c.BodyParser(&payload)
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	err = r.authService.SignUp(ctx, payload)
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return r.ResponseJson(c,
		http.StatusOK,
		nil,
		"Successfully sign up",
		domain.Metadata{},
	)
}

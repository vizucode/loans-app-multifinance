package errorhandler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"multifinancetest/helpers/constants/rpcstd"

	"github.com/vizucode/gokit/utils/errorkit"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type stdRespError struct {
	StatusCode string `json:"status_code"`
	Message    string `json:"message"`
}

// ErrorHandler handles errors and returns an appropriate response.
func FiberErrHandler(ctx *fiber.Ctx, err error) error {

	// Check if the error is a validation error
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		for _, val := range validationErr {
			switch strings.ToLower(val.Tag()) {
			case "required":
				return ctx.Status(fiber.StatusBadRequest).JSON(stdRespError{
					StatusCode: "400" + rpcstd.INVALID_ARGUMENT,
					Message:    strings.ToLower(val.Field()) + " is required",
				})
			case "validate_user_contact":
				return ctx.Status(fiber.StatusBadRequest).JSON(stdRespError{
					StatusCode: "400" + rpcstd.INVALID_ARGUMENT,
					Message:    strings.ToLower(val.Field()) + " must be a valid email address or phone number",
				})
			case "password_regex_validator":
				return ctx.Status(fiber.StatusBadRequest).JSON(stdRespError{
					StatusCode: "400" + rpcstd.INVALID_ARGUMENT,
					Message:    strings.ToLower(val.Field()) + " character must at least 1 uppercase letter, 1 lowercase letter, have a number and 1 special character",
				})
			case "min":
				return ctx.Status(fiber.StatusBadRequest).JSON(stdRespError{
					StatusCode: "400" + rpcstd.INVALID_ARGUMENT,
					Message:    strings.ToLower(val.Field()) + " must be at least " + val.Param() + " characters long",
				})
			case "indonesia_phone":
				return ctx.Status(fiber.StatusBadRequest).JSON(stdRespError{
					StatusCode: "400" + rpcstd.INVALID_ARGUMENT,
					Message:    fmt.Sprintf("%s is not a valid Indonesian phone number format, please enter a valid Indonesian phone number", strings.ToLower(val.Field())),
				})
			case "email":
				return ctx.Status(fiber.StatusBadRequest).JSON(stdRespError{
					StatusCode: "400" + rpcstd.INVALID_ARGUMENT,
					Message:    strings.ToLower("email address is not valid, please enter a valid email address"),
				})
			case "max":
				return ctx.Status(fiber.StatusBadRequest).JSON(stdRespError{
					StatusCode: "400" + rpcstd.INVALID_ARGUMENT,
					Message:    fmt.Sprintf("the length of %s must be at most %s characters", strings.ToLower(val.Field()), val.Param()),
				})
			}
		}
	}

	// check std erro
	var stdError *errorkit.ErrorStd
	ok := errors.As(err, &stdError)
	if ok {
		return ctx.Status(stdError.HttpStatusCode).JSON(stdRespError{
			StatusCode: stdError.ErrorCode(),
			Message:    strings.ToLower(stdError.Error()),
		})
	}

	if err != nil {
		// translate
		return ctx.Status(http.StatusNotFound).JSON(stdRespError{
			StatusCode: "404" + rpcstd.NOT_FOUND,
			Message:    "data was not found",
		})
	}

	return nil
}

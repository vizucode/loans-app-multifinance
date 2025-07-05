package errorhandler

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/vizucode/gokit/utils/errorkit"
)

// ErrorHandler handles errors and returns an appropriate response.
func RpcErrorHandler(err error) (HttpStatusCode int, RpcStatusCode string, message string) {

	// Check if the error is a validation error
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		for _, val := range validationErr {
			switch strings.ToLower(val.Tag()) {
			case "required":
				return 400, "15", strings.ToLower(val.Field()) + " is required"
			case "min":
				return 400, "15", strings.ToLower(val.Field()) + " minimum value is " + strings.ToLower(val.Param())
			}
		}
	}

	// check std erro
	var stdError *errorkit.ErrorStd
	ok := errors.As(err, &stdError)
	if ok {
		return stdError.HttpStatusCode, stdError.RpcStatusCode, stdError.Message
	}

	return
}

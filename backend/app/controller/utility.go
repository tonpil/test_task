package controller

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	prjerror "test_task/interactor/error"
)

func TransformErrorToHttpError(err error) (int, string) {
	switch {
	case errors.Is(err, &prjerror.NotFoundError{}):
		return http.StatusNotFound, "Not found"
	default:
		fmt.Println(err)
		debug.PrintStack()
		return http.StatusInternalServerError, "Internal server error"
	}
}

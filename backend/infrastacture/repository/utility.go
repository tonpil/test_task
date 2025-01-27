package repository

import (
	prjerror "test_task/interactor/error"

	"gopkg.in/reform.v1"
)

func TransformError(err error) error {
	if err == reform.ErrNoRows {
		return &prjerror.NotFoundError{}
	} else {
		return err
	}
}

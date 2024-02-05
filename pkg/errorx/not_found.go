package errorx

import "errors"

type NotFoundError struct {
	BaseError
}

func NewNotFoundError(err error) error {
	return NotFoundError{
		BaseError: NewBaseError(err, ErrorCodeNotFound),
	}
}

func IsNotFoundError(err error) bool {
	return errors.As(err, new(NotFoundError))
}

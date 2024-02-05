package errorx

import "fmt"

type BaseError struct {
	wrapped error
	code    ErrorCode
}

func NewBaseError(err error, code ErrorCode) BaseError {
	return BaseError{
		wrapped: err,
		code:    code,
	}
}

func (b BaseError) Error() string {
	return fmt.Sprintf("%s (%s) ", b.wrapped, b.code)
}

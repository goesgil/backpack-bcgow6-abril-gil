package customError

import (
	"fmt"
)

type CustomError struct {
	msg string
}

func (e *CustomError) InternalError(err error) error {
	return fmt.Errorf("error: internal error, msg: %v", err)
}

func (e *CustomError) BadRequest(err error) error {
	return fmt.Errorf("error: bad request, msg: %v", err)
}

func (e * CustomError) Detail(msg string) error {
	return e.InternalError(fmt.Errorf(msg))
}

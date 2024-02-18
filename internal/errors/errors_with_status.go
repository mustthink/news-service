package errors

import "fmt"

type ErrWithStatus struct {
	status int
	err    error
}

func (e ErrWithStatus) Error() string {
	return e.err.Error()
}

func (e ErrWithStatus) Status() int {
	return e.status
}

func NewErrWithStatus(status int, err error) error {
	return &ErrWithStatus{
		status: status,
		err:    err,
	}
}

func NewErrWithStatusf(status int, format string, a ...interface{}) error {
	return NewErrWithStatus(
		status,
		fmt.Errorf(format, a...),
	)
}

package derror

import "errors"

type DError struct {
	ErrorCode string
	Err       error
}

func New(message string, errCode string) error {
	return &DError{
		ErrorCode: errCode,
		Err:       errors.New(message),
	}
}

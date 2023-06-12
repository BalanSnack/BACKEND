package error

import "errors"

var (
	ErrInvalidRequest                    = errors.New("invalid request")
	ErrInvalidToken                      = errors.New("invalid token")
	ErrInvalidState                      = errors.New("invalid state")
	ErrInvalidResponseFromExternalServer = errors.New("invalid response from external server")
)

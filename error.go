package go_ernie

import (
	"fmt"
)

type APIError struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	ID        string `json:"id"`
}

type RequestError struct {
	HTTPStatusCode int
	Err            error
}

func (e *APIError) Error() string {
	return e.ErrorMsg
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("error, status code: %d, message: %s", e.HTTPStatusCode, e.Err)
}

func (e *RequestError) Unwrap() error {
	return e.Err
}

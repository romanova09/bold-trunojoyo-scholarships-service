package helpers

import "fmt"

type Error struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

var (
	ErrInternal   = Error{Code: "00000", Message: "we encountered an error while processing your request (internal server error)"}
	ErrBadRequest = Error{Code: "00001", Message: "your request is in a bad format"}
)

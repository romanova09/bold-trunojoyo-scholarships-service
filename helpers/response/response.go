package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/romanova09/bold-trunojoyo-scholarship-api/helpers"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

type ErrorResponse struct {
	Success   bool   `json:"success" example:"false"`
	Message   string `json:"message"`
	ErrorCode string `json:"error_code,omitempty"`
	Internal  error  `json:"-"`
}

type BaseResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data"`
}

func buildResponseMsg(defaultMsg string, msg ...string) string {
	if len(msg) == 0 {
		return defaultMsg
	}
	var response string
	for i, item := range msg {
		response += item
		if len(msg)-1 != i {
			response += ", "
		}
	}
	return response
}
func ErrInternalServerError(c *gin.Context, err error) {

	if _, ok := err.(stackTracer); !ok {
		err = errors.WithStack(err)
	}

	originalErr := errors.Cause(err)
	var errorCode string
	var errorMessage string

	if val, ok := originalErr.(helpers.Error); ok {
		errorCode = val.Code
		errorMessage = val.Message
	} else {
		errorCode = helpers.ErrInternal.Code
		errorMessage = helpers.ErrInternal.Message
	}

	resp := ErrorResponse{
		Message:   errorMessage,
		ErrorCode: errorCode,
		Internal:  originalErr,
	}
	Response(c, http.StatusBadRequest, resp, "Internal Server Error")
}

func BadRequest(c *gin.Context, err error) {

	if _, ok := err.(stackTracer); !ok {
		err = errors.WithStack(err)
	}

	originalErr := errors.Cause(err)
	var errorCode string
	var errorMessage string

	if val, ok := originalErr.(helpers.Error); ok {
		errorCode = val.Code
		errorMessage = val.Message
	} else {
		errorCode = helpers.ErrBadRequest.Code
		errorMessage = helpers.ErrBadRequest.Message
	}

	resp := ErrorResponse{
		Message:   errorMessage,
		ErrorCode: errorCode,
		Internal:  originalErr,
	}
	Response(c, http.StatusBadRequest, resp, "Bad Request")
}

func Success(c *gin.Context, data interface{}, msg ...string) {
	responseMsg := buildResponseMsg("Success", msg...)
	Response(c, http.StatusOK, data, responseMsg)
}

func Response(c *gin.Context, code int, data interface{}, msg string) {

	if data == nil {
		data = map[string]interface{}{}
	}
	success := true
	if code > 399 {
		success = false
	}
	res := BaseResponse{
		Success: success,
		Message: msg,
		Data:    data,
	}
	c.JSON(code, res)
}

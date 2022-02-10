package err

import (
	"net/http"
)

// ApiError 自定义api错误结构体
type ApiError struct {
	Status  int         `json:"-"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	//Request string      `json:"request"`
}

func (err ApiError) Error() string {
	return err.Message
}

func newApiError(status int, message string) ApiError {
	return ApiError{
		Status:  status,
		Message: message,
	}
}

func ParamError(err error, message string) ApiError {
	e := newApiError(http.StatusUnprocessableEntity, message)
	e.Data = RequestTranslate(err)
	return e
}

func BadRequestError(err error, message string) ApiError {
	e := newApiError(http.StatusBadRequest, message)
	return e
}

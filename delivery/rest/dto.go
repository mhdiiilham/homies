package rest

import "net/http"

var (
	RequestIDKey = "requestid"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   any    `json:"errors"`
}

func NewSuccessResponse(statusCode int, Data any) Response {
	return Response{Code: statusCode, Message: http.StatusText(statusCode), Data: Data, Error: nil}
}

func NewErroResponse(statusCode int, errs ...error) Response {
	return Response{Code: statusCode, Message: http.StatusText(statusCode), Data: nil, Error: errs}
}

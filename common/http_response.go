package common

import "net/http"

type SuccessResponse struct {
	Code         int         `json:"code"`
	ErrorCode    string      `json:"error_code"`
	ErrorMessage string      `json:"error_message"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}

type ErrorResponse struct {
	Code         int    `json:"code"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Message      string `json:"message"`
	Data         string `json:"data"`
}

func NewErrorResponse(httpCode int, errorMessage error) *ErrorResponse {
	return &ErrorResponse{
		Code:         httpCode,
		ErrorCode:    "ERROR",
		ErrorMessage: errorMessage.Error(),
		Message:      "",
		Data:         "error!",
	}
}

func NewSuccessResponse(data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Code:         http.StatusOK,
		ErrorCode:    "NO_ERROR",
		ErrorMessage: "",
		Message:      "",
		Data:         data,
	}
}

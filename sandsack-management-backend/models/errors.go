package models


type ErrorResponse struct {
	ErrCode    int    `json:"err_code"`
	ErrMessage string `json:"err_message"`
}

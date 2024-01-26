package util

import "readmate/delivery/http/dto/response"

func ResponseError(statusCode int, errorMsg string) (response.ResponseDTO, int) {
	resp := response.ResponseDTO{
		Success:    false,
		StatusCode: statusCode,
		Message:    errorMsg,
	}

	return resp, statusCode
}
func ResponseSuccess(statusCode int, errorMsg string, body any) (response.ResponseDTO, int) {
	resp := response.ResponseDTO{
		Success:    true,
		StatusCode: statusCode,
		Message:    errorMsg,
		Body:       body,
	}

	return resp, statusCode
}

package helpers

import (
	CommonModel "kanbersky/common/models"
	"net/http"
)

func SetHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func PrepareErrorResponse(errorReason string, statusCode int) *CommonModel.ErrorResponse {
	errorResponse := CommonModel.ErrorResponse{}
	errorResponse.Status = statusCode
	errorResponse.Message = errorReason
	return &errorResponse
}

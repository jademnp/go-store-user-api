package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func BadRequestError() *RestErr {
	return &RestErr{
		Message: "invalid json body",
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

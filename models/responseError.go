package models

type ResponseError struct {
	Message string `json:"message"`
	Status  int    `json:"-"`
}

func NewResponseError(message string, status int) *ResponseError {
	return &ResponseError{
		Message: message,
		Status:  status,
	}
}

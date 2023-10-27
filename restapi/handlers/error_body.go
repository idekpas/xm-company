package handlers

type ErrorBody struct {
	ErrorMessage string `json:"error-message"`
}

func (e ErrorBody) Error() string {
	return e.ErrorMessage
}

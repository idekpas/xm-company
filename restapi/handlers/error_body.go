package handlers

type ErrorBody struct {
	ErrorMessage string `json:"error_message"`
}

func (e ErrorBody) Error() string {
	return e.ErrorMessage
}

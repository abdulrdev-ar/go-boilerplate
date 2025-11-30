package exception

type HttpException struct {
	Message    string
	StatusCode int
}

func (e *HttpException) Error() string {
	return e.Message
}

func NewHttpException(message string, statusCode int) *HttpException {
	return &HttpException{
		Message:    message,
		StatusCode: statusCode,
	}
}

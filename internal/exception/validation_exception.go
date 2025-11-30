package exception

func NewValidationException(message string) *HttpException {
	return NewHttpException(message, 400)
}

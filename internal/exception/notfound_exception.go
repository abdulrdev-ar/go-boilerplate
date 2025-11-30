package exception

func NewNotFoundException(message string) *HttpException {
	return NewHttpException(message, 404)
}

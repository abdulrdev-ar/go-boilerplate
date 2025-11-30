package exception

func NewInternalException(message string) *HttpException {
	return NewHttpException(message, 500)
}

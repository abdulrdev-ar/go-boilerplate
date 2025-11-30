package exception

func NewUnauthorizedException(message string) *HttpException {
	return NewHttpException(message, 401)
}

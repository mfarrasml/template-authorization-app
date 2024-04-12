package apperror

type AppError struct {
	err  string
	code int
}

func (e *AppError) Error() string {
	return e.err
}

func (e *AppError) Code() int {
	return e.code
}

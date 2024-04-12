package apperror

var (
	ErrInternalServer = &AppError{
		err:  "internal server error",
		code: 1,
	}

	ErrEmailNotFound = &AppError{
		err:  "email not found",
		code: 2,
	}

	ErrWrongPassword = &AppError{
		err:  "wrong password",
		code: 3,
	}

	ErrAccessToken = &AppError{
		err:  "error generating access token",
		code: 4,
	}
)

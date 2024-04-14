package apperror

var (
	ErrInternalServer = &AppError{
		err:  "internal server error",
		code: 1,
	}

	ErrUserNotFound = &AppError{
		err:  "user not found",
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

	ErrRefreshTokenNotFound = &AppError{
		err:  "refresh token not found",
		code: 5,
	}

	ErrInvalidRefreshToken = &AppError{
		err:  "invalid refresh token",
		code: 6,
	}

	ErrRefreshToken = &AppError{
		err:  "error generating refresh token",
		code: 7,
	}

	ErrParsingRefreshToken = &AppError{
		err:  "error parsing refresh token",
		code: 8,
	}

	ErrInvalidUserId = &AppError{
		err:  "invalid user id",
		code: 10,
	}

	ErrParsingAccessToken = &AppError{
		err:  "error parsing access token",
		code: 20,
	}

	ErrInvalidAccessToken = &AppError{
		err:  "invalid access token",
		code: 21,
	}

	ErrNoRoute = &AppError{
		err:  "route not found",
		code: 40,
	}

	ErrNoMethod = &AppError{
		err:  "method not allowed",
		code: 41,
	}

	ErrBadRequest = &AppError{
		err:  "bad request",
		code: 42,
	}

	ErrForbidden = &AppError{
		err:  "user don't have access to this route",
		code: 43,
	}
)

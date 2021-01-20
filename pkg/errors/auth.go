package errors

type InvalidTokenError struct {Err error}

func (e InvalidTokenError) Error() string {
	return "The token is invalid and cannot serve the purpose of authentication"
}
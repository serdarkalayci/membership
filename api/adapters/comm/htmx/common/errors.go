package common

// ErrCookieNotFound is returned when the login cookie is not found.
type ErrCookieNotFound struct{}

func (e ErrCookieNotFound) Error() string {
	return "login cookie could not be found."
}

// ErrTokenInvalid is returned when the login token is invalid.
type ErrTokenInvalid struct{}

func (e ErrTokenInvalid) Error() string {
	return "login token is invalid."
}

// ErrTokenSignatureInvalid is returned when the login token signature is invalid.
type ErrTokenSignatureInvalid struct{}

func (e ErrTokenSignatureInvalid) Error() string {
	return "token signature is invalid."
}
package auth

import (
	"net/http"
	"strings"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrNoAuthHeaderIncluded Error = "no authorization header included"
	ErrMalformedAuthHeader  Error = "malformed authorization header"
)

// GetAPIKey -
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", ErrMalformedAuthHeader
	}

	return splitAuth[1], nil
}

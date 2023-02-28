package utils

import "errors"

var (
	ErrRecordNotFound = errors.New("Record not found")
	ErrUnauthorized   = errors.New("Unauthorized")
)

package messages

import "errors"

var (
	ErrIncorrectImage = errors.New("Incorrect file type or no user present")
	ErrFailedType     = errors.New("Failed to parse to type")
)

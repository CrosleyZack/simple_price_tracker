package model

import "errors"

var (
	ErrNotFound        = errors.New("Not found")
	ErrEmpty           = errors.New("Empty")
	ErrUnmarshal       = errors.New("Error during Unmarshal()")
	ErrFileOpen        = errors.New("Error when opening file")
	ErrWebsiteNotFound = errors.New("Website not found")
	ErrMarshal         = errors.New("Error during Marshal()")
	ErrFileWrite       = errors.New("Error when writing to file")
)

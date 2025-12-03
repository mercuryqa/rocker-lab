package model

import "errors"

var (
	ErrNotFound            = errors.New("part not found")
	ErrPartListEmpty error = errors.New("part list is empty")
)

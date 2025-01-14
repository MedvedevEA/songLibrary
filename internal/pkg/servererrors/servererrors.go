package servererrors

import "errors"

var (
	ErrorInternal       = errors.New("internal server error")
	ErrorRecordNotFound = errors.New("record not found")
)

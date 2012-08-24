package main

import "fmt"

type hiveError struct {
	ErrStr string
}

func (e *hiveError) Error() string {
	return fmt.Sprintf("%s", e.ErrStr)
}

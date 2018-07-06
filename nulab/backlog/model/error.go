package model

import (
	"bytes"
)

type (
	// Errors represents aggregation of BacklogError.
	Errors struct {
		Errors []Error `json:"errors"`
	}
	// Error represents backlog error.
	Error struct {
		Message  string `json:"message"`
		Code     int    `json:"code"`
		MoreInfo string `json:"moreInfo"`
	}
)

// Error returns error string, implemented error interface
func (e *Errors) Error() string {
	b := bytes.NewBuffer(make([]byte, 0, 1024))
	for i, err := range e.Errors {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(err.Message)
	}
	return b.String()
}

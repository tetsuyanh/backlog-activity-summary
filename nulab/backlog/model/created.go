package model

import "time"

type (
	// CreatedItem represents object knows own created date.
	CreatedItem interface {
		CreatedAt() time.Time
	}
)

type (
	Created struct {
		Created time.Time `json:"created"`
	}
)

func (c *Created) CreatedAt() time.Time {
	return c.Created
}

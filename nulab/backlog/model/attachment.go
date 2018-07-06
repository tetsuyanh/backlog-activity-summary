package model

type (
	// Attachment represents attachment
	Attachment struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Size int    `json:"size"`
	}
)

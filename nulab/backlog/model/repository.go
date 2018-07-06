package model

type (
	// Repository represents repository
	Repository struct {
		ID                 int    `json:"id"`
		Name               string `json:"name"`
		DescriptionSummary string `json:"description"`
	}
)

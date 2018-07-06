package model

type (
	// Issue represents issue
	Issue struct {
		ID          int    `json:"id"`
		KeyID       int    `json:"key_id"`
		Summary     string `json:"summary"`
		Description string `json:"description"`
	}
)

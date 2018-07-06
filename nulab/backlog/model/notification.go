package model

type (
	// Notification represents notification
	Notification struct {
		ID                  int  `json:"id"`
		AlreadyRead         bool `jso:"alreadyRead"`
		Reason              int  `json:"reason"`
		User                User `json:"user"`
		ResourceAlreadyRead bool `json:"resourceAlreadyRead"`
	}
)

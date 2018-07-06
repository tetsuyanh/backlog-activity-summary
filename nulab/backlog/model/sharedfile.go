package model

import "time"

type (
	// SharedFile represents sharedFile
	SharedFile struct {
		ID          int       `json:"id"`
		Type        string    `json:"file"`
		Dir         string    `json:"dir"`
		Name        string    `json:"name"`
		Size        string    `json:"size"`
		CreatedUser User      `json:"createdUser"`
		UpdatedUser User      `json:"updatedUser"`
		Created     time.Time `json:"created"`
		Updated     time.Time `json:"updated"`
	}
)

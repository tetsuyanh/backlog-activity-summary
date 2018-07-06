package model

type (
	// Star represents star
	Star struct {
		Created
		ID        int    `json:"id"`
		Comment   string `json:"comment"`
		URL       string `json:"url"`
		Title     string `json:"title"`
		Presenter User   `json:"presenter"`
	}
)

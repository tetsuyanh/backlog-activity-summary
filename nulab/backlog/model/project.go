package model

type (
	// Project represents project
	Project struct {
		ID                 int    `json:"id"`
		ProjectKey         string `json:"projectKey"`
		Name               string `json:"name"`
		ChartEnabled       bool   `json:"chartEnabled"`
		ProjectLeaderCan   bool   `json:"projectLeaderCanEditProjectLeader"`
		UseWikiTreeView    bool   `json:"useWikiTreeView"`
		TextFormattingRule string `json:"textFormattingRule"`
		Archived           bool   `json:"archived"`
		DisplayOrder       int    `json:"displayOrder"`
	}
)

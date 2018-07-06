package renderer

import (
	"github.com/tetsuyanh/backlog-activity-summary/model"
)

type (
	// Renderer represents to render summary
	Renderer interface {
		RenderSummary(*model.Summary) error
	}
)

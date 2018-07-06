package renderer

import (
	"os"

	"github.com/alecthomas/template"

	"github.com/tetsuyanh/backlog-activity-summary/model"
)

type (
	// MdRenderer is renderer as markdown
	MdRenderer struct{}
)

// NewMdRenderer returns MdRenderer
func NewMdRenderer() Renderer {
	return &MdRenderer{}
}

// RenderSummary renders summary
func (r *MdRenderer) RenderSummary(sm *model.Summary) error {
	tmpl, errParse := template.ParseFiles("template/summary_md.tmpl")
	if errParse != nil {
		return errParse
	}

	f, errCreate := os.Create(`summary.md`)
	if errCreate != nil {
		return errCreate
	}
	defer f.Close()

	errExec := tmpl.Execute(f, sm)
	if errExec != nil {
		return errExec
	}

	return nil
}

package backlog

import "strconv"

const (
	orderAsc  = "asc"
	orderDesc = "desc"
)

type (
	// Params represents parameters of request
	Params struct {
		MinID int
		MaxID int
		Count int
		Order string
	}
)

// NewParams return Params
func NewParams() *Params {
	return &Params{
		MinID: 0,
		MaxID: 0,
		Count: 100,
		Order: orderDesc,
	}
}

func (p *Params) Map() map[string]string {
	m := make(map[string]string, 4)
	if p.MinID > 0 {
		m["minId"] = strconv.Itoa(p.MinID)
	}
	if p.MaxID > 0 {
		m["maxId"] = strconv.Itoa(p.MaxID)
	}
	if p.Count > 0 {
		m["count"] = strconv.Itoa(p.Count)
	}
	if p.Order != orderAsc || p.Order == orderDesc {
		m["order"] = p.Order
	}
	return m
}

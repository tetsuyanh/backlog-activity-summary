package model

import (
	"github.com/tetsuyanh/backlog-activity-summary/nulab/backlog/model"
)

const (
	// SummaryDateFormat is format of summary date
	SummaryDateFormat = "2006/01/02"
)

type (
	// Summary represents various Contributions for period.
	Summary struct {
		Name            string
		Begin           string
		End             string
		ContributionSet map[string]*ContributionSet
	}
)

// NewSummary returns Summary
func NewSummary(name, begin, end string) *Summary {
	return &Summary{
		Name:            name,
		Begin:           begin,
		End:             end,
		ContributionSet: make(map[string]*ContributionSet, 8),
	}
}

// ImportActivities imports activities
func (s *Summary) ImportActivities(acts []*model.Activity) error {
	for _, act := range acts {
		set := s.getContributionSet(act.TypeKind())
		ctrb := set.GetContribution(act.Content.ID, act.Content.Summary)
		switch act.TypeValue() {
		case model.TypeValueCreate:
			ctrb.IncrementCreate()
		case model.TypeValueUpdate:
			ctrb.IncrementUpdate()
		case model.TypeValueDelete:
			ctrb.IncrementDelete()
		case model.TypeValueAction:
			ctrb.IncrementAction()
		}
	}
	return nil
}

// ImportStars imports stars
func (s *Summary) ImportStars(stars []*model.Star) error {
	// TODO
	return nil
}

func (s *Summary) getContributionSet(key string) *ContributionSet {
	set, ok := s.ContributionSet[key]
	if !ok {
		set = NewContributionSet()
		s.ContributionSet[key] = set
	}
	return set
}

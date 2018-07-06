package service

import (
	"fmt"
	"reflect"
	"time"

	"github.com/tetsuyanh/backlog-activity-summary/model"
	"github.com/tetsuyanh/backlog-activity-summary/nulab/backlog"
	bmodel "github.com/tetsuyanh/backlog-activity-summary/nulab/backlog/model"
)

type (
	// UserService represents service around user
	UserService struct {
		bc *backlog.Client
	}
)

// NewUserService returns UserService
func NewUserService() *UserService {
	return &UserService{
		bc: backlog.NewClient(),
	}
}

// Myself returns User
func (s *UserService) Myself() (*bmodel.User, error) {
	return s.bc.Myself()
}

// SummaryPeriod returns summary for a period
func (s *UserService) SummaryPeriod(userID int, name string, begin, end time.Time) (*model.Summary, error) {
	sm := model.NewSummary(
		name,
		begin.Format(model.SummaryDateFormat),
		end.Format(model.SummaryDateFormat),
	)

	cActs := make(chan []*bmodel.Activity)
	defer close(cActs)
	cStars := make(chan []*bmodel.Star)
	defer close(cStars)
	cErr := make(chan error)
	defer close(cErr)

	go func() {
		if acts, errActs := s.ActivitiesPeriod(userID, begin, end); errActs != nil {
			cErr <- errActs
		} else {
			cActs <- acts
		}
	}()
	go func() {
		if stars, errStars := s.StarsPeriod(userID, begin, end); errStars != nil {
			cErr <- errStars
		} else {
			cStars <- stars
		}
	}()

	select {
	case acts := <-cActs:
		stars := <-cStars
		if errActs := sm.ImportActivities(acts); errActs != nil {
			return nil, errActs
		}
		if errStars := sm.ImportStars(stars); errStars != nil {
			return nil, errStars
		}
	case err := <-cErr:
		return nil, err
	}
	return sm, nil
}

// ActivitiesPeriod returns activities for a period
func (s *UserService) ActivitiesPeriod(userID int, begin, end time.Time) ([]*bmodel.Activity, error) {
	p := backlog.NewParams()
	// collect by order desc
	results := make([]*bmodel.Activity, 0, 100)
	for {
		ss, errStars := s.bc.RecentlyActivities(userID, p)
		if errStars != nil {
			return nil, errStars
		}
		h, b, f, errGet := getCreatedRange(ss, begin, end)
		if errGet != nil {
			return nil, errGet
		}
		if h != -1 && b != -1 {
			results = append(results, ss[h:b]...)
		}

		if f || len(ss) < p.Count {
			break
		}
		// specify next range
		p.MaxID = ss[len(ss)-1].ID
	}
	return results, nil
}

// StarsPeriod returns stars for a period
func (s *UserService) StarsPeriod(userID int, begin, end time.Time) ([]*bmodel.Star, error) {
	p := backlog.NewParams()
	// collect by order desc
	results := make([]*bmodel.Star, 0, 100)
	for {
		ss, errStars := s.bc.Stars(userID, p)
		if errStars != nil {
			return nil, errStars
		}
		h, b, f, errGet := getCreatedRange(ss, begin, end)
		if errGet != nil {
			return nil, errGet
		}
		if h != -1 && b != -1 {
			results = append(results, ss[h:b]...)
		}

		if f || len(ss) < p.Count {
			break
		}
		// specify next range
		p.MaxID = ss[len(ss)-1].ID
	}
	return results, nil
}

// getCreatedRange returns range of 'Created' slice between 'begin' and 'end'
// also returns finish flag means no target later
// also returns error of reflection method
// caution) expect slice order by desc
func getCreatedRange(sli interface{}, begin, end time.Time) (int, int, bool, error) {
	val := reflect.ValueOf(sli)
	if val.Kind() != reflect.Slice {
		return -1, -1, true, fmt.Errorf("not kind Array: %s", val.Kind())
	}

	iHead := -1
	iBottom := -1
	noTarget := false
	for i := 0; i < val.Len(); i++ {
		c, ok := val.Index(i).Interface().(bmodel.CreatedItem)
		if !ok {
			return -1, -1, true, fmt.Errorf("not type Created")
		}

		t := c.CreatedAt()
		if t.After(end) {
			continue
		}
		if t.Before(begin) {
			noTarget = true
			break
		}
		if iHead == -1 {
			iHead = i
		}
		iBottom = i + 1
	}
	return iHead, iBottom, noTarget, nil
}

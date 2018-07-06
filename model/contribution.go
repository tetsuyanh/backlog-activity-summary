package model

type (
	// Contribution represents contribution
	Contribution struct {
		Subject string
		Create  int
		Update  int
		Delete  int
		Action  int
		Stared  int
	}

	// ContributionSet represents contribution
	ContributionSet struct {
		ContributionMap map[int]*Contribution
	}
)

// NewContributionSet returns NewContribution
func NewContributionSet() *ContributionSet {
	return &ContributionSet{
		ContributionMap: make(map[int]*Contribution, 16),
	}
}

// GetContribution returns GetContribution
func (c *ContributionSet) GetContribution(key int, subject string) *Contribution {
	ctrb, ok := c.ContributionMap[key]
	if !ok {
		ctrb = NewContribution(subject)
		c.ContributionMap[key] = ctrb
	}
	return ctrb
}

// NewContribution returns Contribution
func NewContribution(subject string) *Contribution {
	return &Contribution{
		Subject: subject,
	}
}

// IncrementCreate increments count of Create
func (c *Contribution) IncrementCreate() {
	c.Create++
}

// IncrementUpdate increments count of Update
func (c *Contribution) IncrementUpdate() {
	c.Update++
}

// IncrementDelete increments count of Delete
func (c *Contribution) IncrementDelete() {
	c.Delete++
}

// IncrementAction increments count of Action
func (c *Contribution) IncrementAction() {
	c.Action++
}

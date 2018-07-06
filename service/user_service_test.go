package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/tetsuyanh/backlog-activity-summary/nulab/backlog/model"
)

var (
	b   = time.Date(2018, 6, 1, 0, 0, 0, 0, time.UTC)
	e   = time.Date(2018, 6, 30, 23, 59, 59, 999, time.UTC)
	sBf = &model.Star{Created: model.Created{Created: time.Date(2018, 5, 31, 23, 59, 59, 999, time.UTC)}}
	sIn = &model.Star{Created: model.Created{Created: time.Date(2018, 6, 15, 0, 0, 0, 0, time.UTC)}}
	sAf = &model.Star{Created: model.Created{Created: time.Date(2018, 7, 1, 0, 0, 0, 0, time.UTC)}}
)

type (
	Case struct {
		s       interface{}
		b       time.Time
		e       time.Time
		eHead   int
		eBottom int
		eFlag   bool
		eErr    bool
		label   string
	}
)

func TestGetCreatedRange(t *testing.T) {
	cases := []Case{
		// normal case
		Case{[]model.CreatedItem{sAf, sAf, sAf}, b, e, -1, -1, false, false, "[after]"},
		Case{[]model.CreatedItem{sAf, sAf, sIn}, b, e, 2, 3, false, false, "[after, in]"},
		Case{[]model.CreatedItem{sAf, sAf, sBf}, b, e, -1, -1, true, false, "[after, before]"},
		Case{[]model.CreatedItem{sAf, sIn, sBf}, b, e, 1, 2, true, false, "[after, in, before]"},
		Case{[]model.CreatedItem{sIn, sIn, sIn}, b, e, 0, 3, false, false, "[in]"},
		Case{[]model.CreatedItem{sIn, sIn, sBf}, b, e, 0, 2, true, false, "[in, before]"},
		Case{[]model.CreatedItem{sBf, sBf, sBf}, b, e, -1, -1, true, false, "[before]"},
		// error case
		Case{[]int{1, 2, 3}, b, e, -1, -1, true, true, "int slice"},
		Case{[]interface{}{sAf, sAf, 3}, b, e, -1, -1, true, true, "include type int"},
	}
	for _, v := range cases {
		h, b, f, e := getCreatedRange(v.s, v.b, v.e)
		assert.Equal(t, v.eHead, h, v.label)
		assert.Equal(t, v.eBottom, b, v.label)
		assert.Equal(t, v.eFlag, f, v.label)
		if v.eErr {
			assert.NotNil(t, e, v.label)
		} else {
			assert.Nil(t, e, v.label)
		}
	}
}

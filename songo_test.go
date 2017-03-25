package songo

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

var (
	testURL1 string = "http://127.0.0.1/demo" +
		"?_limit=50&_page=2" +
		"&_sort=created,money,-level" +
		"&year=$eq$2016&month=$bt$8,11&date=$eq$1&day=$in$0,6" +
		"&time=$eq$201612182359"
)

func TestSongo_ParseRawURL(t *testing.T) {
	var s Songo
	s.Include("year")
	s.Exclude("time")
	if err := s.ParseRawURL(testURL1); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, s.Limit, 50)
	assert.Equal(t, s.Page, 2)
	assert.Equal(t, s.Sort, []string([]string{"created", "money", "-level"}))
	if os, v, ok := s.Query.GetQuery("year"); ok {
		assert.Equal(t, os, []string{"$eq"})
		assert.Equal(t, v, int64(2016))
	}
	if os, v, ok := s.Query.GetQuery("month"); ok {
		assert.Equal(t, os, []string{"$bt"})
		assert.Equal(t, v, []string([]string{"8", "11"}))
	}
	if os, v, ok := s.Query.GetQuery("date"); ok {
		assert.Equal(t, os, []string{"$eq"})
		assert.Equal(t, v, int64(1))
	}
	if os, v, ok := s.Query.GetQuery("day"); ok {
		assert.Equal(t, os, []string{"$in"})
		assert.Equal(t, v, []string{"0", "6"})
	}
	if _, _, ok := s.Query.GetQuery("time"); ok {
		assert.FailNow(t, "Fail to exclude key", "time")
	}
}

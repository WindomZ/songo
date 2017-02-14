package songo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSongo_ParseRawURL(t *testing.T) {
	var s Songo
	if err := s.ParseRawURL("http://127.0.0.1/demo" +
		"?_limit=50&_page=2" +
		"&_sort=created,money,-level" +
		"&year=$eq$2016&month=$bt$8,11&date=$eq$1&day=$in$0,6"); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, s.Limit, 50)
	assert.Equal(t, s.Page, 2)
	assert.Equal(t, s.Sort, []string([]string{"created", "money", "-level"}))
	if k, v := s.Query.GetQuery("year"); v != nil {
		assert.Equal(t, k, "$eq")
		assert.Equal(t, v, int64(2016))
	}
	if k, v := s.Query.GetQuery("month"); v != nil {
		assert.Equal(t, k, "$bt")
		assert.Equal(t, v, []string([]string{"8", "11"}))
	}
	if k, v := s.Query.GetQuery("date"); v != nil {
		assert.Equal(t, k, "$eq")
		assert.Equal(t, v, true)
	}
	if k, v := s.Query.GetQuery("day"); v != nil {
		assert.Equal(t, k, "$in")
		assert.Equal(t, v, []string{"0", "6"})
	}
}

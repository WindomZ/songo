package mongo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testURL1 string = "http://127.0.0.1/demo" +
		"?_limit=50&_page=2" +
		"&_sort=created,money,-level" +
		"&year=$eq$2016&month=$bt$8,11&date=$eq$1&day=$in$0,6" +
		"&time=$or$gt$100&time=$or$lt$200"

	testURL2 string = "http://127.0.0.1/demo" +
		"?_limit=50&_page=2" +
		"&_sort=created,money,-level" +
		"&$or$time=$gt$100, $lt$200"
)

func TestSongo_Result1(t *testing.T) {
	var s Songo
	if err := s.ParseRawURL(testURL1); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, s.Limit, 50)
	assert.Equal(t, s.Page, 2)
	assert.Equal(t, s.Sort, []string([]string{"created", "money", "-level"}))

	r := s.Result()
	if v, ok := r["year"]; ok {
		assert.Equal(t, v, int64(2016))
	} else {
		assert.FailNow(t, "Should exist 'year'")
	}
	if v, ok := r["month"]; ok {
		assert.Equal(t, v, SongoResultMap{"$bt": []string{"8", "11"}})
	} else {
		assert.FailNow(t, "Should exist 'month'")
	}
	if v, ok := r["date"]; ok {
		assert.Equal(t, v, int64(1))
	} else {
		assert.FailNow(t, "Should exist 'date'")
	}
	if v, ok := r["day"]; ok {
		assert.Equal(t, v, SongoResultMap{"$in": []string{"0", "6"}})
	} else {
		assert.FailNow(t, "Should exist 'day'")
	}
	if v, ok := r["$or"]; ok {
		assert.Equal(t, v, []SongoResultMap{
			{"$gt": int64(100)},
			{"$lt": int64(200)},
		})
	} else {
		assert.FailNow(t, "Should exist '$or'")
	}
	if v, ok := r["time"]; ok {
		assert.FailNow(t, fmt.Sprintf("Should not exist: %#v", v))
	}
}

func TestSongo_Result2(t *testing.T) {
	var s Songo
	if err := s.ParseRawURL(testURL2); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, s.Limit, 50)
	assert.Equal(t, s.Page, 2)
	assert.Equal(t, s.Sort, []string([]string{"created", "money", "-level"}))

	r := s.Result()
	if v, ok := r["$or"]; ok {
		assert.Equal(t, v, []SongoResultMap{
			{"$gt": int64(100)},
			{"$lt": int64(200)},
		})
	} else {
		assert.FailNow(t, "Should exist '$or'")
	}
	if v, ok := r["time"]; ok {
		assert.FailNow(t, fmt.Sprintf("Should not exist: %#v", v))
	}
}

package songo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var query SongoQuery

func TestSongoQuery_Init(t *testing.T) {
	query.Set("$or$and$int", "$eq$123")
	query.Set("$or$and$bool", "$eq$true")
	query.Set("$and$or$string", "$eq$hello world")
	query.Set("$and$or$string", "$eq$how are u")
}

func TestSongoQuery_GetKeys(t *testing.T) {
	keys := query.GetKeys()
	assert.Equal(t, keys, []string{"int", "bool", "string"})
}

func TestSongoQuery_GetValues(t *testing.T) {
	if values, ok := query.GetValues("int"); ok {
		assert.Equal(t, len(values), 1)
		assert.Equal(t, values[0].GetValue(), int64(123))
	} else {
		assert.Panics(t, nil)
	}
	if values, ok := query.GetValues("bool"); ok {
		assert.Equal(t, len(values), 1)
		assert.Equal(t, values[0].GetValue(), true)
	} else {
		assert.Panics(t, nil)
	}
	if values, ok := query.GetValues("string"); ok {
		assert.Equal(t, len(values), 2)
		assert.Equal(t, values[0].GetValue(), "hello world")
		assert.Equal(t, values[1].GetValue(), "how are u")
	} else {
		assert.Panics(t, nil)
	}
}

func TestSongoQuery_GetQuery(t *testing.T) {
	if os, v, ok := query.GetQuery("int"); ok {
		assert.Equal(t, os, []string([]string{"$or", "$and", "$eq"}))
		assert.Equal(t, v, int64(123))
	}
	if os, v, ok := query.GetQuery("bool"); ok {
		assert.Equal(t, os, []string([]string{"$or", "$and", "$eq"}))
		assert.Equal(t, v, true)
	}
	if os, v, ok := query.GetQuery("string", 0); ok {
		assert.Equal(t, os, []string([]string{"$and", "$or", "$eq"}))
		assert.Equal(t, v, "hello world")
	}
	if os, v, ok := query.GetQuery("string", 1); ok {
		assert.Equal(t, os, []string([]string{"$and", "$or", "$eq"}))
		assert.Equal(t, v, "how are u")
	}
}

package songo

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestVerifyQueryValue(t *testing.T) {
	assert.Equal(t, VerifyQueryValue("$eq$xxx"), true)
	assert.Equal(t, VerifyQueryValue("$eq1$xxx"), false)

	assert.Equal(t, VerifyQueryValue("$and$eq$xxx"), true)
	assert.Equal(t, VerifyQueryValue("$and1$eq$xxx"), false)
	assert.Equal(t, VerifyQueryValue("$and$eq2$xxx"), false)
}

func TestSongoQueryValue_ValueStrings(t *testing.T) {
	s, ok := SplitQueryValue("$and$eq$xxx")

	assert.Equal(t, ok, true)
	assert.Equal(t, s.ValueStrings(), []string{"$and", "$eq", "xxx"})
}

package songo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifyQueryValue(t *testing.T) {
	assert.Equal(t, VerifyQueryValue("$eq$xxx"), true)
	assert.Equal(t, VerifyQueryValue("$eq1$xxx"), false)
}

package ex05_test

import (
	"gpl/ch4/ex05"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduct(t *testing.T) {
	s := []int{1, 2, 2, 3, 5, 6, 6}
	act, err := ex05.Reduct(s)
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3, 5, 6}, act)
}

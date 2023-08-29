package ex01_test

import (
	"crypto/sha256"
	"gpl/ch4/ex01"
	"testing"
)

func TestCheckbit(t *testing.T) {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	diff := ex01.Checkbit(c1, c2)
	if diff != 125 {
		t.Errorf("diff: %d", diff)
	}
}

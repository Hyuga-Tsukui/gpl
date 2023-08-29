package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountDiff(x, y [32]byte) int {
	var count int
	for i := 0; i < len(x); i++ {
		count += int(pc[x[i]^y[i]])
	}
	return count
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	fmt.Printf("c1: %x\n", c1)
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("c2: %x\n", c2)

	diff := PopCountDiff(c1, c2)
	fmt.Printf("diff: %d\n", diff)

}

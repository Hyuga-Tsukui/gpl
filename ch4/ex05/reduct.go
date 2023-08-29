package ex05

import "fmt"

func Reduct(x []int) ([]int, error) {
	if len(x) == 0 {
		return nil, fmt.Errorf("empty slice")
	}

	var c int
	var s []int

	for _, v := range x {
		if c == v {
			continue
		}
		s = append(s, v)
		c = v
	}
	return s, nil
}

package ex01

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func Checkbit(x, y [32]byte) int {
	var count int
	for i := 0; i < len(x); i++ {
		count += int(pc[x[i]^y[i]])
	}
	return count
}

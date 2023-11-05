package ch4

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func Embed() {
	w := Wheel{Circle{Point{8, 8}, 5}, 20}
	w.X = 42
}

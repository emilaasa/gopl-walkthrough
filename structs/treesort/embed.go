package embed

type Point struct {
	X, Y int
}

type Cricle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

var w = Wheel{Circle{Point{8, 8}, 5}, 20}

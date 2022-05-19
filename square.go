package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (square Square) End() Point {
	return Point{square.start.x + int(square.a), square.start.y - int(square.a)}
}

func (square Square) Area() uint {
	return square.a * square.a
}

func (square Square) Perimeter() uint {
	return 4 * square.a
}

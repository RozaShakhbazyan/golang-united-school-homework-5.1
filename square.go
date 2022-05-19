package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (square Square) End() Point {
	var x int = square.start.x
	var y int = square.start.y
	if x > 0 {
		x = x + int(square.a)
	} else {
		x = x - int(square.a)
	}
	if y > 0 {
		y = y + int(square.a)
	} else {
		y = y - int(square.a)
	}
	return Point{x, y}
}

func (square Square) Area() uint {
	return square.a * square.a
}

func (square Square) Perimeter() uint {
	return 4 * square.a
}

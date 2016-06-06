package shape

type Circle struct {
	R int64
}

func (c *Circle) area() int64 {
	return c.R*c.R
}

type Square struct {
	Width int64
	Height int64
}

func (s *Square) area() int64 {
	return s.Width*s.Height
}

type Shape interface {
	area() int64
}

func ShapeArea(s Shape) int64 {
	return s.area()
}

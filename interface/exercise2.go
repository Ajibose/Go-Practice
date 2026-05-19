package inter

import (
	"math"
)

type Geometry interface {
	Area() float64
	Perim() float64
}

type Rect struct {
	Width, Height float64
}
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rect) Area() float64 {
	return r.Width * r.Height
}
func (r Rect) Perim() float64 {
	return 2*r.Width + 2*r.Height
}


// # AlertThreshold represent the rectangle being worked on.
//
// it is a representation of.
// Though it is seen as ngfje.
// “[JSON and Go].”
// "[RFC 7159]."
//
// [RFC 7159]: https://tools.ietf.org/html/rfc7159
// [JSON and Go]: https://golang.org/doc/articles/json_and_go.html
func AlertThreshold(g Geometry) string {
	if c, ok := g.(Circle); ok && c.Radius > 10 {
		return "Circle is too large!"
	} else if r, ok := g.(Rect); ok && r.Area() > 100 {
		return "Rectangle is too large!"
	} else {
		return "Shape is within limits."
	}
}

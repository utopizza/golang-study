package ch6

import (
	"math"
)

// Point 点
type Point struct {
	X, Y float64
}

// Distance 距离
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance 距离
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// ScaleBy 扩展
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

//-----------------------------------------------//

// Path 路径
type Path []Point

// Distance 距离
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

//-----------------------------------------------//

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	// method value
	distanceFromP := p.Distance
	distanceFromP(q) //p.Distance(q)
	scaleP := p.ScaleBy
	scaleP(2) //p.ScaleBy(2)

	// method expression(the first param is receiver)
	distance := Point.Distance
	distance(p, q) //p.Distance(q)
	scale := (*Point).ScaleBy
	scale(&p, 2) //p.ScaleBy(2)
}

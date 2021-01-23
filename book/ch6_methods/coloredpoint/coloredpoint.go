package ch6

import (
	"image/color"

	geometry "golang-study/book/ch6_methods/geometry"
)

// ColoredPoint 组合
type ColoredPoint struct {
	geometry.Point
	Color color.RGBA
}

// Distance 距离
func (p ColoredPoint) Distance(q geometry.Point) float64 {
	return p.Point.Distance(q)
}

// ScaleBy 扩展
func (p *ColoredPoint) ScaleBy(factor float64) {
	p.Point.ScaleBy(factor)
}

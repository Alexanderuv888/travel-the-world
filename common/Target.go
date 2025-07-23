package common

import "image"

type Target interface {
	Point() image.Point
}

type TargetFunc func() image.Point

func (f TargetFunc) Point() image.Point {
	return f()
}

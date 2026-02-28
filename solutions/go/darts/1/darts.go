package darts

import (
	"math"
)

func Score(x, y float64) int {
	toDart := math.Sqrt(x*x + y*y)
	switch {
	case toDart <= 1:
		return 10
	case toDart <= 5:
		return 5
	case toDart <= 10:
		return 1
	default:
		return 0
	}
}

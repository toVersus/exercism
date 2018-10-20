// Package triangle classifies triangles by calculating side-lengths.
package triangle

import (
	"math"
	"sort"
)

// Kind represents types of triangle
type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides classify triangle based on the input lengths of sides.
func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}
	sort.Float64s(sides)

	switch {
	case !isTriangle(sides):
		return NaT
	case isEquilateral(sides):
		return Equ
	case issIsosceles(sides):
		return Iso
	default:
		return Sca
	}
}

func isTriangle(sorted []float64) bool {
	return !math.IsNaN(sorted[0]) && sorted[0] > 0 && !math.IsInf(sorted[2], 1) &&
		sorted[0]+sorted[1] >= sorted[2]
}

func isEquilateral(sorted []float64) bool {
	return sorted[0] == sorted[2]
}

func issIsosceles(sorted []float64) bool {
	return sorted[0] == sorted[1] || sorted[1] == sorted[2]
}

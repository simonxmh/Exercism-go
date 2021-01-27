// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package space

// import (
// 	"math"
// )
type Planet string


// Age calculates the age of a person on a certain planet
func Age(seconds float64, p Planet) float64 {
	table := map[Planet]float64{
		"Earth": seconds/31557600.00, 
		"Mercury": seconds/31557600.00 / 0.2408467,
		"Venus": seconds/31557600.00 / 0.61519726,
		"Mars": seconds/31557600.00 / 1.8808158,
		"Jupiter": seconds/31557600.00 / 11.862615,
		"Saturn": seconds/31557600.00 / 29.447498,
		"Uranus": seconds/31557600.00 / 84.016846,
		"Neptune": seconds/31557600.00 / 164.79132,
	}
	return table[p]
}

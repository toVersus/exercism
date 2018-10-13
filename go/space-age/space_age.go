package space

import (
	"math"
)

// Planet represents planet name
type Planet string

const (
	EARTH_YEAR   = 365.25 // day
	EARTH_DAY    = 24     // hours
	EARTH_HOUR   = 60     // minutes
	EARTH_MINUTE = 60     //seconds

	EARTH_AGE   = 31557600   // Earth seconds
	MERCURY_AGE = 0.2408467  // Earth years
	VENUS_AGE   = 0.61519726 // Earth years
	MARS_AGE    = 1.8808158  // Earth years
	JUPITER_AGE = 11.862615  // Earth years
	SATURN_AGE  = 29.447498  // Earth years
	URANUS_AGE  = 84.016846  // Earth years
	NEPTUNE_AGE = 164.79132  // Earth years
)

var ageList = map[Planet]float64{
	"Earth":   EARTH_AGE,
	"Mercury": earthYearsToSeconds(MERCURY_AGE),
	"Venus":   earthYearsToSeconds(VENUS_AGE),
	"Mars":    earthYearsToSeconds(MARS_AGE),
	"Jupiter": earthYearsToSeconds(JUPITER_AGE),
	"Saturn":  earthYearsToSeconds(SATURN_AGE),
	"Uranus":  earthYearsToSeconds(URANUS_AGE),
	"Neptune": earthYearsToSeconds(NEPTUNE_AGE),
}

func earthYearsToSeconds(years float64) float64 {
	return years * EARTH_YEAR * EARTH_DAY * EARTH_HOUR * EARTH_MINUTE
}

func round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Round(f*shift) / shift
}

// Age computes someone's age in terms of its planet time scale.
func Age(seconds float64, planet Planet) float64 {
	return round(seconds/ageList[planet], 2)
}

package space

import "math"

type Planet string

const secondsCountInEarthYear = 31557600

var relativeToEarthYear = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func roundTo2Decimals(f float64) float64 {
	return math.Round(f*100) / 100
}

func Age(sc float64, p Planet) float64 {
	inEarth := sc / secondsCountInEarthYear
	inPlanet := inEarth / relativeToEarthYear[p]

	return roundTo2Decimals(inPlanet)
}

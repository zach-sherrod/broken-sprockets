package angle

import (
	"math"

	"github.com/zach-sherrod/broken_sprockets/quantity"
)

const (
	radiansPerDegree    = math.Pi / 180.0
	degreesPerRadian    = 180.0 / math.Pi
	degreesPerArcMinute = 1.0 / 60.0
	degreesPerArcSecond = 1.0 / 3600.0
)

func Radians(v float64) quantity.Angle {
	return quantity.NewAngleFromRadians(v)
}

func Degrees(v float64) quantity.Angle {
	return quantity.NewAngleFromRadians(v * radiansPerDegree)
}

func ArcMinutes(v float64) quantity.Angle {
	return Degrees(v * degreesPerArcMinute)
}

func ArcSeconds(v float64) quantity.Angle {
	return Degrees(v * degreesPerArcSecond)
}

func ToRadians(a quantity.Angle) float64 {
	return a.Radians()
}

func ToDegrees(a quantity.Angle) float64 {
	return a.Radians() * degreesPerRadian
}

func ToArcMinutes(a quantity.Angle) float64 {
	return ToDegrees(a) * 60.0
}

func ToArcSeconds(a quantity.Angle) float64 {
	return ToDegrees(a) * 3600.0
}

func NormalizeZeroToTwoPi(a quantity.Angle) quantity.Angle {
	r := math.Mod(a.Radians(), 2.0*math.Pi)
	if r < 0 {
		r += 2.0 * math.Pi
	}
	return quantity.NewAngleFromRadians(r)
}

func NormalizeMinusPiToPi(a quantity.Angle) quantity.Angle {
	r := math.Mod(a.Radians()+math.Pi, 2.0*math.Pi)
	if r < 0 {
		r += 2.0 * math.Pi
	}
	r -= math.Pi
	return quantity.NewAngleFromRadians(r)
}

func NormalizeZeroTo360(a quantity.Angle) quantity.Angle {
	return Degrees(ToDegrees(NormalizeZeroToTwoPi(a)))
}

func NormalizeMinus180To180(a quantity.Angle) quantity.Angle {
	return Degrees(ToDegrees(NormalizeMinusPiToPi(a)))
}

// SignedDelta returns the shortest signed angular difference from "from" to "to".
// Result is normalized to (-π, π].
func SignedDelta(from, to quantity.Angle) quantity.Angle {
	return NormalizeMinusPiToPi(to.Sub(from))
}

// AbsDelta returns the smallest absolute angular separation between two angles.
func AbsDelta(a1, a2 quantity.Angle) quantity.Angle {
	d := SignedDelta(a1, a2)
	r := d.Radians()
	if r < 0 {
		r = -r
	}
	return quantity.NewAngleFromRadians(r)
}

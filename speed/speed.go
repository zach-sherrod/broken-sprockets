package speed

import (
	"github.com/zach-sherrod/broken_sprockets/distance"
	"github.com/zach-sherrod/broken_sprockets/duration"
	"github.com/zach-sherrod/broken_sprockets/quantity"
)

const (
	metersPerSecondPerFootPerSecond = 0.3048
	metersPerSecondPerKilometerHour = 1000.0 / 3600.0
	metersPerSecondPerMileHour      = 1609.344 / 3600.0
	metersPerSecondPerKnot          = 1852.0 / 3600.0
)

func MetersPerSecond(v float64) quantity.Speed {
	return quantity.NewSpeedFromMetersPerSecond(v)
}

func FeetPerSecond(v float64) quantity.Speed {
	return quantity.NewSpeedFromMetersPerSecond(v * metersPerSecondPerFootPerSecond)
}

func KilometersPerHour(v float64) quantity.Speed {
	return quantity.NewSpeedFromMetersPerSecond(v * metersPerSecondPerKilometerHour)
}

func MilesPerHour(v float64) quantity.Speed {
	return quantity.NewSpeedFromMetersPerSecond(v * metersPerSecondPerMileHour)
}

func Knots(v float64) quantity.Speed {
	return quantity.NewSpeedFromMetersPerSecond(v * metersPerSecondPerKnot)
}

func ToMetersPerSecond(s quantity.Speed) float64 {
	return s.MetersPerSecond()
}

func ToFeetPerSecond(s quantity.Speed) float64 {
	return s.MetersPerSecond() / metersPerSecondPerFootPerSecond
}

func ToKilometersPerHour(s quantity.Speed) float64 {
	return s.MetersPerSecond() / metersPerSecondPerKilometerHour
}

func ToMilesPerHour(s quantity.Speed) float64 {
	return s.MetersPerSecond() / metersPerSecondPerMileHour
}

func ToKnots(s quantity.Speed) float64 {
	return s.MetersPerSecond() / metersPerSecondPerKnot
}

func FromDistanceOverDuration(d quantity.Distance, t quantity.Duration) quantity.Speed {
	return quantity.NewSpeedFromMetersPerSecond(
		distance.ToMeters(d) / duration.ToSeconds(t),
	)
}

func ToDistanceOverDuration(s quantity.Speed, t quantity.Duration) quantity.Distance {
	return distance.Meters(
		s.MetersPerSecond() * duration.ToSeconds(t),
	)
}

package acceleration

import (
	"github.com/zach-sherrod/broken_sprockets/duration"
	"github.com/zach-sherrod/broken_sprockets/quantity"
	"github.com/zach-sherrod/broken_sprockets/speed"

	"errors"
)

const (
	metersPerSecondSquaredPerFootPerSecondSquared = 0.3048
	standardGravityMetersPerSecondSquared         = 9.80665
)

var ErrZeroDuration = errors.New("duration must not be zero")

func MetersPerSecondSquared(v float64) quantity.Acceleration {
	return quantity.NewAccelerationFromMetersPerSecondSquared(v)
}

func FeetPerSecondSquared(v float64) quantity.Acceleration {
	return quantity.NewAccelerationFromMetersPerSecondSquared(v * metersPerSecondSquaredPerFootPerSecondSquared)
}

func G(v float64) quantity.Acceleration {
	return quantity.NewAccelerationFromMetersPerSecondSquared(v * standardGravityMetersPerSecondSquared)
}

func ToMetersPerSecondSquared(a quantity.Acceleration) float64 {
	return a.MetersPerSecondSquared()
}

func ToFeetPerSecondSquared(a quantity.Acceleration) float64 {
	return a.MetersPerSecondSquared() / metersPerSecondSquaredPerFootPerSecondSquared
}

func ToG(a quantity.Acceleration) float64 {
	return a.MetersPerSecondSquared() / standardGravityMetersPerSecondSquared
}

func FromSpeedOverDuration(s quantity.Speed, d quantity.Duration) quantity.Acceleration {
	return quantity.NewAccelerationFromMetersPerSecondSquared(
		speed.ToMetersPerSecond(s) / duration.ToSeconds(d),
	)
}

func TryFromSpeedOverDuration(s quantity.Speed, d quantity.Duration) (quantity.Acceleration, error) {
	if d.IsZero() {
		return quantity.Acceleration{}, ErrZeroDuration
	}

	return quantity.NewAccelerationFromMetersPerSecondSquared(
		speed.ToMetersPerSecond(s) / duration.ToSeconds(d),
	), nil
}

func ToSpeedOverDuration(a quantity.Acceleration, d quantity.Duration) quantity.Speed {
	return speed.MetersPerSecond(
		a.MetersPerSecondSquared() * duration.ToSeconds(d),
	)
}

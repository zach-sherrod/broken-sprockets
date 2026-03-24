package angularvelocity

import (
	"errors"
	"math"

	"github.com/zach-sherrod/broken-sprockets/angle"
	"github.com/zach-sherrod/broken-sprockets/duration"
	"github.com/zach-sherrod/broken-sprockets/quantity"
)

const (
	radiansPerSecondPerDegreePerSecond    = math.Pi / 180.0
	radiansPerSecondPerArcMinutePerSecond = radiansPerSecondPerDegreePerSecond / 60.0
	radiansPerSecondPerArcSecondPerSecond = radiansPerSecondPerDegreePerSecond / 3600.0
)

var ErrZeroDuration = errors.New("duration must not be zero")

func RadiansPerSecond(v float64) quantity.AngularVelocity {
	return quantity.NewAngularVelocityFromRadiansPerSecond(v)
}

func DegreesPerSecond(v float64) quantity.AngularVelocity {
	return quantity.NewAngularVelocityFromRadiansPerSecond(v * radiansPerSecondPerDegreePerSecond)
}

func ArcMinutesPerSecond(v float64) quantity.AngularVelocity {
	return quantity.NewAngularVelocityFromRadiansPerSecond(v * radiansPerSecondPerArcMinutePerSecond)
}

func ArcSecondsPerSecond(v float64) quantity.AngularVelocity {
	return quantity.NewAngularVelocityFromRadiansPerSecond(v * radiansPerSecondPerArcSecondPerSecond)
}

func ToRadiansPerSecond(w quantity.AngularVelocity) float64 {
	return w.RadiansPerSecond()
}

func ToDegreesPerSecond(w quantity.AngularVelocity) float64 {
	return angle.ToDegrees(angle.Radians(w.RadiansPerSecond()))
}

func ToArcMinutesPerSecond(w quantity.AngularVelocity) float64 {
	return ToDegreesPerSecond(w) * 60.0
}

func ToArcSecondsPerSecond(w quantity.AngularVelocity) float64 {
	return ToDegreesPerSecond(w) * 3600.0
}

func FromAngleOverDuration(a quantity.Angle, d quantity.Duration) quantity.AngularVelocity {
	return quantity.NewAngularVelocityFromRadiansPerSecond(
		angle.ToRadians(a) / duration.ToSeconds(d),
	)
}

func TryFromAngleOverDuration(a quantity.Angle, d quantity.Duration) (quantity.AngularVelocity, error) {
	if d.IsZero() {
		return quantity.AngularVelocity{}, ErrZeroDuration
	}

	return quantity.NewAngularVelocityFromRadiansPerSecond(
		angle.ToRadians(a) / duration.ToSeconds(d),
	), nil
}

func ToAngleOverDuration(w quantity.AngularVelocity, d quantity.Duration) quantity.Angle {
	return angle.Radians(
		w.RadiansPerSecond() * duration.ToSeconds(d),
	)
}

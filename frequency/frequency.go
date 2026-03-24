package frequency

import (
	"errors"

	"github.com/zach-sherrod/broken-sprockets/duration"
	"github.com/zach-sherrod/broken-sprockets/quantity"
)

const (
	hertzPerKilohertz = 1_000.0
	hertzPerMegahertz = 1_000_000.0
	hertzPerGigahertz = 1_000_000_000.0
	rpmPerHertz       = 60.0
)

var (
	ErrZeroDuration  = errors.New("duration must not be zero")
	ErrZeroFrequency = errors.New("frequency must not be zero")
)

func Hertz(v float64) quantity.Frequency {
	return quantity.NewFrequencyFromHertz(v)
}

func Kilohertz(v float64) quantity.Frequency {
	return quantity.NewFrequencyFromHertz(v * hertzPerKilohertz)
}

func Megahertz(v float64) quantity.Frequency {
	return quantity.NewFrequencyFromHertz(v * hertzPerMegahertz)
}

func Gigahertz(v float64) quantity.Frequency {
	return quantity.NewFrequencyFromHertz(v * hertzPerGigahertz)
}

func RPM(v float64) quantity.Frequency {
	return quantity.NewFrequencyFromHertz(v / rpmPerHertz)
}

func ToHertz(f quantity.Frequency) float64 {
	return f.Hertz()
}

func ToKilohertz(f quantity.Frequency) float64 {
	return f.Hertz() / hertzPerKilohertz
}

func ToMegahertz(f quantity.Frequency) float64 {
	return f.Hertz() / hertzPerMegahertz
}

func ToGigahertz(f quantity.Frequency) float64 {
	return f.Hertz() / hertzPerGigahertz
}

func ToRPM(f quantity.Frequency) float64 {
	return f.Hertz() * rpmPerHertz
}

func FromPeriod(d quantity.Duration) quantity.Frequency {
	return quantity.NewFrequencyFromHertz(1.0 / duration.ToSeconds(d))
}

func TryFromPeriod(d quantity.Duration) (quantity.Frequency, error) {
	if d.IsZero() {
		return quantity.Frequency{}, ErrZeroDuration
	}

	return quantity.NewFrequencyFromHertz(1.0 / duration.ToSeconds(d)), nil
}

func ToPeriod(f quantity.Frequency) quantity.Duration {
	return duration.Seconds(1.0 / f.Hertz())
}

func TryToPeriod(f quantity.Frequency) (quantity.Duration, error) {
	if f.IsZero() {
		return quantity.Duration{}, ErrZeroFrequency
	}

	return duration.Seconds(1.0 / f.Hertz()), nil
}

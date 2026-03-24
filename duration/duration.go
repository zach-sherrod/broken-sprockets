package duration

import "github.com/zach-sherrod/broken_sprockets/quantity"

const (
	nanosecondsPerMicrosecond = int64(1_000)
	nanosecondsPerMillisecond = int64(1_000_000)
	nanosecondsPerSecond      = int64(1_000_000_000)
	nanosecondsPerMinute      = int64(60) * nanosecondsPerSecond
	nanosecondsPerHour        = int64(60) * nanosecondsPerMinute
)

func Nanoseconds(v int64) quantity.Duration {
	return quantity.NewDurationFromNanoseconds(v)
}

func Microseconds(v int64) quantity.Duration {
	return quantity.NewDurationFromNanoseconds(v * nanosecondsPerMicrosecond)
}

func Milliseconds(v int64) quantity.Duration {
	return quantity.NewDurationFromNanoseconds(v * nanosecondsPerMillisecond)
}

func Seconds(v float64) quantity.Duration {
	return quantity.NewDurationFromNanoseconds(int64(v * float64(nanosecondsPerSecond)))
}

func Minutes(v float64) quantity.Duration {
	return quantity.NewDurationFromNanoseconds(int64(v * float64(nanosecondsPerMinute)))
}

func Hours(v float64) quantity.Duration {
	return quantity.NewDurationFromNanoseconds(int64(v * float64(nanosecondsPerHour)))
}

func ToNanoseconds(d quantity.Duration) int64 {
	return d.Nanoseconds()
}

func ToMicroseconds(d quantity.Duration) int64 {
	return d.Nanoseconds() / nanosecondsPerMicrosecond
}

func ToMilliseconds(d quantity.Duration) int64 {
	return d.Nanoseconds() / nanosecondsPerMillisecond
}

func ToSeconds(d quantity.Duration) float64 {
	return float64(d.Nanoseconds()) / float64(nanosecondsPerSecond)
}

func ToMinutes(d quantity.Duration) float64 {
	return float64(d.Nanoseconds()) / float64(nanosecondsPerMinute)
}

func ToHours(d quantity.Duration) float64 {
	return float64(d.Nanoseconds()) / float64(nanosecondsPerHour)
}

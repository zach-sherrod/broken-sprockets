package quantity

import "time"

type Duration struct {
	nanoseconds int64
}

func NewDurationFromNanoseconds(v int64) Duration {
	return Duration{nanoseconds: v}
}

func (duration Duration) Nanoseconds() int64 {
	return duration.nanoseconds
}

func (duration Duration) Stdlib() time.Duration {
	return time.Duration(duration.nanoseconds)
}

func FromStdlibDuration(v time.Duration) Duration {
	return Duration{nanoseconds: int64(v)}
}

func (duration Duration) Add(duration2 Duration) Duration {
	return Duration{nanoseconds: duration.nanoseconds + duration2.nanoseconds}
}

func (duration Duration) Sub(duration2 Duration) Duration {
	return Duration{nanoseconds: duration.nanoseconds - duration2.nanoseconds}
}

func (duration Duration) MulScalar(s float64) Duration {
	return Duration{nanoseconds: int64(float64(duration.nanoseconds) * s)}
}

func (duration Duration) DivScalar(s float64) Duration {
	return Duration{nanoseconds: int64(float64(duration.nanoseconds) / s)}
}

func (duration Duration) IsZero() bool {
	return duration.nanoseconds == 0
}

func (duration Duration) LessThan(duration2 Duration) bool {
	return duration.nanoseconds < duration2.nanoseconds
}

func (duration Duration) GreaterThan(duration2 Duration) bool {
	return duration.nanoseconds > duration2.nanoseconds
}

package quantity

type Frequency struct {
	hertz float64
}

func NewFrequencyFromHertz(v float64) Frequency {
	return Frequency{hertz: v}
}

func (frequency Frequency) Hertz() float64 {
	return frequency.hertz
}

func (frequency Frequency) Add(frequency2 Frequency) Frequency {
	return Frequency{hertz: frequency.hertz + frequency2.hertz}
}

func (frequency Frequency) Sub(frequency2 Frequency) Frequency {
	return Frequency{hertz: frequency.hertz - frequency2.hertz}
}

func (frequency Frequency) MulScalar(k float64) Frequency {
	return Frequency{hertz: frequency.hertz * k}
}

func (frequency Frequency) DivScalar(k float64) Frequency {
	return Frequency{hertz: frequency.hertz / k}
}

func (frequency Frequency) IsZero() bool {
	return frequency.hertz == 0
}

func (frequency Frequency) LessThan(frequency2 Frequency) bool {
	return frequency.hertz < frequency2.hertz
}

func (frequency Frequency) GreaterThan(frequency2 Frequency) bool {
	return frequency.hertz > frequency2.hertz
}

func (frequency Frequency) EqualWithin(frequency2 Frequency, tolerance Frequency) bool {
	diff := frequency.hertz - frequency2.hertz
	if diff < 0 {
		diff = -diff
	}
	return diff <= tolerance.hertz
}

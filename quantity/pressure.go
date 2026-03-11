package quantity

type Pressure struct {
	pascals float64
}

func NewPressureFromPascals(v float64) Pressure {
	return Pressure{pascals: v}
}

func (pressure Pressure) Pascals() float64 {
	return pressure.pascals
}

func (pressure Pressure) Add(pressure2 Pressure) Pressure {
	return Pressure{pascals: pressure.pascals + pressure2.pascals}
}

func (pressure Pressure) Sub(pressure2 Pressure) Pressure {
	return Pressure{pascals: pressure.pascals - pressure2.pascals}
}

func (pressure Pressure) MulScalar(k float64) Pressure {
	return Pressure{pascals: pressure.pascals * k}
}

func (pressure Pressure) DivScalar(k float64) Pressure {
	return Pressure{pascals: pressure.pascals / k}
}

func (pressure Pressure) IsZero() bool {
	return pressure.pascals == 0
}

func (pressure Pressure) LessThan(pressure2 Pressure) bool {
	return pressure.pascals < pressure2.pascals
}

func (pressure Pressure) GreaterThan(pressure2 Pressure) bool {
	return pressure.pascals > pressure2.pascals
}

func (pressure Pressure) EqualWithin(pressure2 Pressure, tolerance Pressure) bool {
	diff := pressure.pascals - pressure2.pascals
	if diff < 0 {
		diff = -diff
	}
	return diff <= tolerance.pascals
}

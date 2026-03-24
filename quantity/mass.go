package quantity

type Mass struct {
	kilograms float64
}

func NewMassFromKilograms(v float64) Mass {
	return Mass{kilograms: v}
}

func (mass Mass) Kilograms() float64 {
	return mass.kilograms
}

func (mass Mass) Add(mass2 Mass) Mass {
	return Mass{kilograms: mass.kilograms + mass2.kilograms}
}

func (mass Mass) Sub(mass2 Mass) Mass {
	return Mass{kilograms: mass.kilograms - mass2.kilograms}
}

func (mass Mass) IsZero() bool {
	return mass.kilograms == 0
}

func (mass Mass) LessThan(mass2 Mass) bool {
	return mass.kilograms < mass2.kilograms
}

func (mass Mass) GreaterThan(mass2 Mass) bool {
	return mass.kilograms > mass2.kilograms
}

func (mass Mass) EqualWithin(mass2 Mass, tolerance Mass) bool {
	diff := mass.kilograms - mass2.kilograms
	if diff < 0 {
		diff = -diff
	}
	return diff <= tolerance.kilograms
}

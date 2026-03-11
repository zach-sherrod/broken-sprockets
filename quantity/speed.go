package quantity

type Speed struct {
	metersPerSecond float64
}

func NewSpeedFromMetersPerSecond(v float64) Speed {
	return Speed{metersPerSecond: v}
}

func (speed Speed) MetersPerSecond() float64 {
	return speed.metersPerSecond
}

func (speed Speed) Add(speed2 Speed) Speed {
	return Speed{metersPerSecond: speed.metersPerSecond + speed2.metersPerSecond}
}

func (speed Speed) Sub(speed2 Speed) Speed {
	return Speed{metersPerSecond: speed.metersPerSecond - speed2.metersPerSecond}
}

func (speed Speed) IsZero() bool {
	return speed.metersPerSecond == 0
}

func (speed Speed) LessThan(speed2 Speed) bool {
	return speed.metersPerSecond < speed2.metersPerSecond
}

func (speed Speed) GreaterThan(speed2 Speed) bool {
	return speed.metersPerSecond > speed2.metersPerSecond
}

func (speed Speed) EqualWithin(speed2 Speed, tolerance Speed) bool {
	diff := speed.metersPerSecond - speed2.metersPerSecond
	if diff < 0 {
		diff = -diff
	}
	return diff <= tolerance.metersPerSecond
}

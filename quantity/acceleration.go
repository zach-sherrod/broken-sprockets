package quantity

type Acceleration struct {
	metersPerSecondSquared float64
}

func NewAccelerationFromMetersPerSecondSquared(v float64) Acceleration {
	return Acceleration{metersPerSecondSquared: v}
}

func (acceleration Acceleration) MetersPerSecondSquared() float64 {
	return acceleration.metersPerSecondSquared
}

func (acceleration Acceleration) Add(acceleration2 Acceleration) Acceleration {
	return Acceleration{metersPerSecondSquared: acceleration.metersPerSecondSquared + acceleration2.metersPerSecondSquared}
}

func (acceleration Acceleration) Sub(acceleration2 Acceleration) Acceleration {
	return Acceleration{metersPerSecondSquared: acceleration.metersPerSecondSquared - acceleration2.metersPerSecondSquared}
}

func (acceleration Acceleration) IsZero() bool {
	return acceleration.metersPerSecondSquared == 0
}

func (acceleration Acceleration) LessThan(acceleration2 Acceleration) bool {
	return acceleration.metersPerSecondSquared < acceleration2.metersPerSecondSquared
}

func (acceleration Acceleration) GreaterThan(acceleration2 Acceleration) bool {
	return acceleration.metersPerSecondSquared > acceleration2.metersPerSecondSquared
}

func (acceleration Acceleration) EqualWithin(acceleration2 Acceleration, tolerance Acceleration) bool {
	diff := acceleration.metersPerSecondSquared - acceleration2.metersPerSecondSquared
	if diff < 0 {
		diff = -diff
	}
	return diff <= tolerance.metersPerSecondSquared
}

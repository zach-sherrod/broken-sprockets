package quantity

type AngularVelocity struct {
	radiansPerSecond float64
}

func NewAngularVelocityFromRadiansPerSecond(v float64) AngularVelocity {
	return AngularVelocity{radiansPerSecond: v}
}

func (angularVelocity AngularVelocity) RadiansPerSecond() float64 {
	return angularVelocity.radiansPerSecond
}

func (angularVelocity AngularVelocity) Add(angularVelocity2 AngularVelocity) AngularVelocity {
	return AngularVelocity{radiansPerSecond: angularVelocity.radiansPerSecond + angularVelocity2.radiansPerSecond}
}

func (angularVelocity AngularVelocity) Sub(angularVelocity2 AngularVelocity) AngularVelocity {
	return AngularVelocity{radiansPerSecond: angularVelocity.radiansPerSecond - angularVelocity2.radiansPerSecond}
}

func (angularVelocity AngularVelocity) IsZero() bool {
	return angularVelocity.radiansPerSecond == 0
}

func (angularVelocity AngularVelocity) LessThan(angularVelocity2 AngularVelocity) bool {
	return angularVelocity.radiansPerSecond < angularVelocity2.radiansPerSecond
}

func (angularVelocity AngularVelocity) GreaterThan(angularVelocity2 AngularVelocity) bool {
	return angularVelocity.radiansPerSecond > angularVelocity2.radiansPerSecond
}

func (angularVelocity AngularVelocity) EqualWithin(angularVelocity2 AngularVelocity, tolerance AngularVelocity) bool {
	diff := angularVelocity.radiansPerSecond - angularVelocity2.radiansPerSecond
	if diff < 0 {
		diff = -diff
	}
	return diff <= tolerance.radiansPerSecond
}

package quantity

type Angle struct {
	radians float64
}

func NewAngleFromRadians(v float64) Angle {
	return Angle{radians: v}
}

func (angle Angle) Radians() float64 {
	return angle.radians
}

func (angle Angle) Add(angle2 Angle) Angle {
	return Angle{radians: angle.radians + angle2.radians}
}

func (angle Angle) Sub(angle2 Angle) Angle {
	return Angle{radians: angle.radians - angle2.radians}
}

func (angle Angle) MulScalar(s float64) Angle {
	return Angle{radians: angle.radians * s}
}

func (angle Angle) DivScalar(s float64) Angle {
	return Angle{radians: angle.radians / s}
}

func (angle Angle) IsZero() bool {
	return angle.radians == 0
}

func (angle Angle) LessThan(angle2 Angle) bool {
	return angle.radians < angle2.radians
}

func (angle Angle) GreaterThan(angle2 Angle) bool {
	return angle.radians > angle2.radians
}

func (angle Angle) EqualWithin(angle2 Angle, tolerance Angle) bool {
	diff := angle.radians - angle2.radians
	if diff < 0 {
		diff = -diff
	}
	return diff <= tolerance.radians
}

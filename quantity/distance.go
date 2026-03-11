package quantity

type Distance struct {
	meters float64
}

func NewDistanceFromMeters(v float64) Distance {
	return Distance{meters: v}
}

func (distance Distance) Meters() float64 {
	return distance.meters
}

func (distance Distance) Add(distance2 Distance) Distance {
	return Distance{meters: distance.meters + distance2.meters}
}

func (distance Distance) Sub(distance2 Distance) Distance {
	return Distance{meters: distance.meters - distance2.meters}
}

func (distance Distance) IsZero() bool {
	return distance.meters == 0
}

func (distance Distance) LessThan(distance2 Distance) bool {
	return distance.meters < distance2.meters
}

func (distance Distance) GreaterThan(distance2 Distance) bool {
	return distance.meters > distance2.meters
}

func (distance Distance) EqualWithin(distance2 Distance, tolerance Distance) bool {
	diff := distance.meters - distance2.meters
	if diff < 0 {
		diff = -diff
	}
	return diff <= tolerance.meters
}

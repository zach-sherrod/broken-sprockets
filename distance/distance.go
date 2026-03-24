package distance

import "github.com/zach-sherrod/broken-sprockets/quantity"

const (
	metersPerKilometer    = 1000.0
	metersPerCentimeter   = 0.01
	metersPerMillimeter   = 0.001
	metersPerInch         = 0.0254
	metersPerFoot         = 0.3048
	metersPerYard         = 0.9144
	metersPerMile         = 1609.344
	metersPerNauticalMile = 1852.0
)

func Meters(v float64) quantity.Distance {
	return quantity.NewDistanceFromMeters(v)
}

func Kilometers(v float64) quantity.Distance {
	return quantity.NewDistanceFromMeters(v * metersPerKilometer)
}

func Centimeters(v float64) quantity.Distance {
	return quantity.NewDistanceFromMeters(v * metersPerCentimeter)
}

func Millimeters(v float64) quantity.Distance {
	return quantity.NewDistanceFromMeters(v * metersPerMillimeter)
}

func Inches(v float64) quantity.Distance {
	return quantity.NewDistanceFromMeters(v * metersPerInch)
}

func Feet(v float64) quantity.Distance {
	return quantity.NewDistanceFromMeters(v * metersPerFoot)
}

func Yards(v float64) quantity.Distance {
	return quantity.NewDistanceFromMeters(v * metersPerYard)
}

func Miles(v float64) quantity.Distance {
	return quantity.NewDistanceFromMeters(v * metersPerMile)
}

func NauticalMiles(v float64) quantity.Distance {
	return quantity.NewDistanceFromMeters(v * metersPerNauticalMile)
}

func ToMeters(d quantity.Distance) float64 {
	return d.Meters()
}

func ToKilometers(d quantity.Distance) float64 {
	return d.Meters() / metersPerKilometer
}

func ToCentimeters(d quantity.Distance) float64 {
	return d.Meters() / metersPerCentimeter
}

func ToMillimeters(d quantity.Distance) float64 {
	return d.Meters() / metersPerMillimeter
}

func ToInches(d quantity.Distance) float64 {
	return d.Meters() / metersPerInch
}

func ToFeet(d quantity.Distance) float64 {
	return d.Meters() / metersPerFoot
}

func ToYards(d quantity.Distance) float64 {
	return d.Meters() / metersPerYard
}

func ToMiles(d quantity.Distance) float64 {
	return d.Meters() / metersPerMile
}

func ToNauticalMiles(d quantity.Distance) float64 {
	return d.Meters() / metersPerNauticalMile
}

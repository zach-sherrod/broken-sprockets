package mass

import "github.com/zach-sherrod/broken_sprockets/quantity"

const (
	kilogramsPerGram      = 0.001
	kilogramsPerMilligram = 0.000001
	kilogramsPerPound     = 0.45359237
	kilogramsPerOunce     = 0.028349523125
	kilogramsPerMetricTon = 1000.0
)

func Kilograms(v float64) quantity.Mass {
	return quantity.NewMassFromKilograms(v)
}

func Grams(v float64) quantity.Mass {
	return quantity.NewMassFromKilograms(v * kilogramsPerGram)
}

func Milligrams(v float64) quantity.Mass {
	return quantity.NewMassFromKilograms(v * kilogramsPerMilligram)
}

func Pounds(v float64) quantity.Mass {
	return quantity.NewMassFromKilograms(v * kilogramsPerPound)
}

func Ounces(v float64) quantity.Mass {
	return quantity.NewMassFromKilograms(v * kilogramsPerOunce)
}

func MetricTons(v float64) quantity.Mass {
	return quantity.NewMassFromKilograms(v * kilogramsPerMetricTon)
}

func ToKilograms(m quantity.Mass) float64 {
	return m.Kilograms()
}

func ToGrams(m quantity.Mass) float64 {
	return m.Kilograms() / kilogramsPerGram
}

func ToMilligrams(m quantity.Mass) float64 {
	return m.Kilograms() / kilogramsPerMilligram
}

func ToPounds(m quantity.Mass) float64 {
	return m.Kilograms() / kilogramsPerPound
}

func ToOunces(m quantity.Mass) float64 {
	return m.Kilograms() / kilogramsPerOunce
}

func ToMetricTons(m quantity.Mass) float64 {
	return m.Kilograms() / kilogramsPerMetricTon
}

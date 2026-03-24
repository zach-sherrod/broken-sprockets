package pressure

import "github.com/zach-sherrod/broken-sprockets/quantity"

const (
	pascalsPerKilopascal = 1_000.0
	pascalsPerMegapascal = 1_000_000.0
	pascalsPerBar        = 100_000.0
	pascalsPerMillibar   = 100.0
	pascalsPerPSI        = 6_894.757293168
	pascalsPerAtmosphere = 101_325.0
	pascalsPerTorr       = 133.32236842105263
)

func Pascals(v float64) quantity.Pressure {
	return quantity.NewPressureFromPascals(v)
}

func Kilopascals(v float64) quantity.Pressure {
	return quantity.NewPressureFromPascals(v * pascalsPerKilopascal)
}

func Megapascals(v float64) quantity.Pressure {
	return quantity.NewPressureFromPascals(v * pascalsPerMegapascal)
}

func Bar(v float64) quantity.Pressure {
	return quantity.NewPressureFromPascals(v * pascalsPerBar)
}

func Millibar(v float64) quantity.Pressure {
	return quantity.NewPressureFromPascals(v * pascalsPerMillibar)
}

func PSI(v float64) quantity.Pressure {
	return quantity.NewPressureFromPascals(v * pascalsPerPSI)
}

func Atmospheres(v float64) quantity.Pressure {
	return quantity.NewPressureFromPascals(v * pascalsPerAtmosphere)
}

func Torr(v float64) quantity.Pressure {
	return quantity.NewPressureFromPascals(v * pascalsPerTorr)
}

func ToPascals(p quantity.Pressure) float64 {
	return p.Pascals()
}

func ToKilopascals(p quantity.Pressure) float64 {
	return p.Pascals() / pascalsPerKilopascal
}

func ToMegapascals(p quantity.Pressure) float64 {
	return p.Pascals() / pascalsPerMegapascal
}

func ToBar(p quantity.Pressure) float64 {
	return p.Pascals() / pascalsPerBar
}

func ToMillibar(p quantity.Pressure) float64 {
	return p.Pascals() / pascalsPerMillibar
}

func ToPSI(p quantity.Pressure) float64 {
	return p.Pascals() / pascalsPerPSI
}

func ToAtmospheres(p quantity.Pressure) float64 {
	return p.Pascals() / pascalsPerAtmosphere
}

func ToTorr(p quantity.Pressure) float64 {
	return p.Pascals() / pascalsPerTorr
}

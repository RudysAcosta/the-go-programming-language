package exercises

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	Freezingc     Celsius = 0
	Boilingc      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FTpC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

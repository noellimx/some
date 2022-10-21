package calculator

type Fareinheit float64
type Kelvin float64
type Celsius float64

type Units interface {
	Fareinheit | Kelvin | Celsius
}

type UnitDefault Fareinheit

type inputUnitTypeFlag int

type conv_foo[I Units, J Units] func(I) J

type temp struct {
	InputType     inputUnitTypeFlag
	TempCelsius   Celsius
	TempFarenheit Fareinheit
	TempKelvin    Kelvin

	ConvertFtoC conv_foo[Fareinheit, Celsius]
	ConvertCtoK conv_foo[Celsius, Kelvin]
	ConvertKtoF conv_foo[Kelvin, Fareinheit]
}

const (
	UK inputUnitTypeFlag = 1
	F
	C
	K
)

func FtoC(x Fareinheit) Celsius {
	return Celsius(5/9*float64(x) - 32)
}

func CtoK(x Celsius) Kelvin {
	return Kelvin(float64(x) + 273.15)
}

func KtoF(x Kelvin) Fareinheit {
	return Fareinheit(1.8*(float64(x)-273) + 32)
}

func CtoF(x Celsius) Fareinheit {
	return Fareinheit(x*1.8000 + 32.0)

}

func strToInputUnitType(i int) inputUnitTypeFlag {

	switch i {
	case 1:
		return F
	case 2:
		return K
	case 3:
		return C
	default:
		return UK
	}

}

func (t *temp) convertAll(i int, v float64) {

	t.InputType = strToInputUnitType(i)

	if t.InputType == F {
		t.TempFarenheit = Fareinheit(v)
	}

}
func NewTemp(i int, v float64) *temp {
	t := &temp{
		ConvertFtoC: FtoC,
		ConvertKtoF: KtoF,
		ConvertCtoK: CtoK,
	}
	if t.InputType == UK {
		return nil
	}

	return t
}

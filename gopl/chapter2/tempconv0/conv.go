package tempconv0

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(k*9/5 - 459.67)
}

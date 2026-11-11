package tempconv

// CToF converts a Celsius temperature to Farenheit.
func CToF(c Celsius) Farenheit { return Farenheit(c*9/5 + 32) }

// FToC converts a Farenheit temperature to Celsuis.
func FToC(f Farenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius((k - 273.15)) }

package smath

// GCD ...
func GCD(a, b int) int {
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}
	if b > a {
		return GCD(b, a)
	}
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

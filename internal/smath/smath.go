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

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

// Abs ...
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

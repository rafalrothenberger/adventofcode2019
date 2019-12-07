package day

// Day4Task2 ...
func Day4Task2() interface{} {
	res := 0
	for i := 124075; i <= 580769; i++ {
		digits := toDigits(i)
		hasDouble := false
		sameNumbers := 1
		isNotDecreasing := true
		for i := 0; i < 5; i++ {
			if digits[i] > digits[i+1] {
				isNotDecreasing = false
				break
			}

			if digits[i] == digits[i+1] {
				sameNumbers++
			} else {
				if sameNumbers == 2 {
					hasDouble = true
				}
				sameNumbers = 1
			}
		}

		if sameNumbers == 2 {
			hasDouble = true
		}

		if isNotDecreasing && hasDouble {
			res++
		}
	}
	return res
}

// Day4Task1 ...
func Day4Task1() interface{} {
	res := 0
	for i := 124075; i <= 580769; i++ {
		digits := toDigits(i)
		hasDouble := false
		isNotDecreasing := true
		for i := 0; i < 5; i++ {
			if digits[i] > digits[i+1] {
				isNotDecreasing = false
				break
			}

			if digits[i] == digits[i+1] {
				hasDouble = true
			}
		}

		if isNotDecreasing && hasDouble {
			res++
		}
	}
	return res
}

func toDigits(v int) []int {
	digits := make([]int, 6)

	for i := 5; i >= 0; i-- {
		digits[i] = v % 10
		v /= 10
	}

	return digits
}

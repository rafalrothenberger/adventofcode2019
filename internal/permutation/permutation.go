package permutation

// Number ...
func Number(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

// Gen ...
func Gen(permutation []int) [][]int {
	return gen([]int{}, permutation)
}

func gen(permutation, rest []int) [][]int {
	if len(rest) > 0 {
		result := make([][]int, 0)
		for _, r := range rest {
			x := append(permutation, r)
			y := filter(rest, r)
			for _, k := range gen(x, y) {
				result = append(result, k)
			}
		}
		return result
	}
	return [][]int{permutation}
}

func filter(list []int, e int) []int {
	res := make([]int, 0)
	for _, l := range list {
		if l == e {
			continue
		}
		res = append(res, l)
	}
	return res
}

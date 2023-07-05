package pascal

func Triangle(n int) [][]int {
	result := make([][]int, n)
	result[0] = []int{1}
	if n == 1 {
		return result
	}
	result[1] = []int{1, 1}
	if n == 2 {
		return result
	}
	for i := 2; i < n; i += 1 {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		for j := 1; j < i; j += 1 {
			result[i][j] = result[i-1][j] + result[i-1][j-1]
		}
		result[i][i] = 1
	}
	return result
}

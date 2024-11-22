package arrays_slices

func Sum(a []int) int {
	var sum int

	for _, x := range a {
		sum += x
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	var sums []int
	for _, theSlice := range slices {
		sums = append(sums, Sum(theSlice))
	}
	return sums
}

func SumAllTails(slices ...[]int) []int {
	var sums []int

	for _, theSlice := range slices {
		if len(theSlice) < 2 {
			sums = append(sums, Sum(theSlice))
			continue
		}

		sums = append(sums, Sum(theSlice[1:]))
	}
	return sums
}

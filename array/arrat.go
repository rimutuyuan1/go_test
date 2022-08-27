package array

func Array(numbers [5]int) int {
	var total int

	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}

	return total
}

package slice

// 数组和切片的区别
// 数组的长度也是类型的一部分，[4]int与[5]int不是同种类型，比较笨重
// 切片类型为[]int，不需要指定切片长度
func slice_sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// make函数是创建切片的新方式，可以指定切片长度和容量
func slice_sum_sum_all(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = slice_sum(numbers)
	}

	return sums
}

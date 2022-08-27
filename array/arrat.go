package array

import "fmt"

//使用for循环累加数组的值
func Array(numbers [5]int) int {
	var total int

	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}

	return total
}

//使用range代替for循环
func Array2(numbers [5]int) int {
	var total int

	for i, num := range numbers {
		total += num
		fmt.Println(i)
	}

	return total
}

//使用range代替for循环
//range每次返回数组元素的索引和值，如果不需要某个参数，可以使用_代替
func Array3(numbers [5]int) int {
	var total int

	for _, num := range numbers {
		total += num
	}

	return total
}

package slice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		//此时 numbers类型为切片（未指定长度）
		numbers := []int{1, 2, 3, 4, 5}

		got := slice_sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := slice_sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {

	got := slice_sum_sum_all([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	//在Go中不能对切片使用等号运算符，可以使用reflect.DeepEqual比较两个变量是否相等
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

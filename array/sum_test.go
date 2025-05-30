package array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("want %d but got %d", want, got)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{0, 4})
	want := []int{6, 4}

	checkSums(t, got, want)
}

func TestSumAllTail(t *testing.T) {
	t.Run("with not empty array", func(t *testing.T) {
		got := SumAllTail([]int{1, 2, 4}, []int{3, 3})
		want := []int{6, 3}

		checkSums(t, got, want)
	})

	t.Run("with empty array", func(t *testing.T) {
		got := SumAllTail([]int{1, 2, 4}, []int{3, 3}, []int{})
		want := []int{6, 3, 0}

		checkSums(t, got, want)
	})
}

func checkSums(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v but got %v", want, got)
	}
}

func ExampleSum() {
	Sumed := Sum([]int{2, 4})
	fmt.Println(Sumed)

	// Output: 6
}

func ExampleSumAll() {
	Result := SumAll([]int{1, 3, 4}, []int{3, 4})
	fmt.Println(Result)

	// Output: [8 7]
}

func ExampleSumAllTail() {
	Result := SumAllTail([]int{2, 3, 5}, []int{4, 6, 2, 2})
	fmt.Println(Result)

	// Output: [8 10]
}

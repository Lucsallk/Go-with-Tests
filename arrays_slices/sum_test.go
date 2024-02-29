package arraysslices

import (
	"testing"
	// "reflect"
	"slices"
)

func TestSum(t *testing.T) {
	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	// want := "bob"
	want := []int{3, 9}

	// Go não deixa com que operadores de igualdade em slices
	// Err: if got != want {
	// reflect.DeepEqual pode ser usado mas não é typeSage
	// if !reflect.DeepEqual(got, want) {
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

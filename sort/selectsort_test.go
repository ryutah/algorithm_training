package sort

import (
	"reflect"
	"testing"
)

func TestSelectSort(t *testing.T) {
	arr := []int{10, 3, 4, 1, 4, 5, 0, 12}
	got := SelectSort(arr)
	want := []int{0, 1, 3, 4, 4, 5, 10, 12}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestSelectSortSortEmptyList(t *testing.T) {
	arr := []int{}
	got := SelectSort(arr)
	want := []int{}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestSelectSortSortedList(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := SelectSort(arr)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestSelectSortSameValueList(t *testing.T) {
	arr := []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	got := SelectSort(arr)
	want := []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func BenchmarkSelectSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectSort(randomArray())
	}
}

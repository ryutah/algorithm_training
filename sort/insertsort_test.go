package sort

import (
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := []int{10, 3, 4, 1, 4, 5, 0, 12}
	got := InsertSort(arr)
	want := []int{0, 1, 3, 4, 4, 5, 10, 12}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
func TestInsertSortSortEmptyList(t *testing.T) {
	arr := []int{}
	got := InsertSort(arr)
	want := []int{}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestInsertSortSortedList(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := InsertSort(arr)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestInsertSortSameValueList(t *testing.T) {
	arr := []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	got := InsertSort(arr)
	want := []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func BenchmarkInsertSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertSort(randomArray())
	}
}

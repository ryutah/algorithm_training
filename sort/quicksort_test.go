package sort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	arr := []int{10, 3, 4, 1, 4, 5, 0, 12, 6, 8}
	got := QuickSort(arr)
	want := []int{0, 1, 3, 4, 4, 5, 6, 8, 10, 12}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestQuickSortSortEmptyList(t *testing.T) {
	arr := []int{}
	got := QuickSort(arr)
	want := []int{}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestQuickSortSortedList(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := QuickSort(arr)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestQuickSortSameValueList(t *testing.T) {
	arr := []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	got := QuickSort(arr)
	want := []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(randomArray())
	}
}

func randomArray() []int {
	rand.Seed(time.Now().UnixNano())
	const len = 100000
	arr := make([]int, len)
	for i := 0; i < len; i++ {
		arr[i] = rand.Int()
	}
	return arr
}

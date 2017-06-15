package sort_test

import (
	"testing"

	. "github.com/ryutah/algorithm_training/sort"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HeapSort", func() {
	Context("default List", func() {
		It("should sort list", func() {
			arr := []int{10, 3, 4, 1, 4, 5, 0, 12}
			want := []int{0, 1, 3, 4, 4, 5, 10, 12}
			got := HeapSort(arr)
			Expect(got).To(Equal(want))
		})

		It("should not sort source list", func() {
			arr := []int{10, 3, 4, 1, 4, 5, 0, 12}
			want := []int{10, 3, 4, 1, 4, 5, 0, 12}
			HeapSort(arr)
			Expect(arr).To(Equal(want))
		})
	})

	Context("empty list", func() {
		It("should not be happend error", func() {
			defer func() {
				err := recover()
				Expect(err).To(BeNil())
			}()
			HeapSort([]int{})
		})
		It("should return empty list", func() {
			got := HeapSort([]int{})
			Expect(got).To(BeEmpty())
		})
	})

	Context("sorted list", func() {
		It("should not be happen error", func() {
			defer func() {
				err := recover()
				Expect(err).To(BeNil())
			}()
			arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			HeapSort(arr)
		})
		It("should be return sort list", func() {
			arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			got := HeapSort(arr)
			Expect(got).To(Equal(want))
		})
	})

	Context("same value list", func() {
		It("should not be happen error", func() {
			defer func() {
				err := recover()
				Expect(err).To(BeNil())
			}()
			arr := []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
			HeapSort(arr)
		})
		It("should return sorted list", func() {
			arr := []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
			want := []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
			got := HeapSort(arr)
			Expect(got).To(Equal(want))
		})
	})
})

func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HeapSort(randomArray())
	}
}
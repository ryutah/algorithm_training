package sort

func MergeSort(arr []int) []int {
	copyarray := make([]int, len(arr))
	copy(copyarray, arr)
	startMergeSort(copyarray)
	return copyarray
}

func startMergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	middle := len(arr) / 2
	remain := len(arr) - middle
	arr1, arr2 := make([]int, middle), make([]int, remain)
	copy(arr1, arr[:middle])
	copy(arr2, arr[middle:])
	startMergeSort(arr1)
	startMergeSort(arr2)
	merge(arr, arr1, arr2)
}

func merge(dst, arr1, arr2 []int) {
	for i, j := 0, 0; i < len(arr1) || j < len(arr2); {
		if j >= len(arr2) || (i < len(arr1) && arr1[i] < arr2[j]) {
			dst[i+j] = arr1[i]
			i++
		} else {
			dst[i+j] = arr2[j]
			j++
		}
	}
}

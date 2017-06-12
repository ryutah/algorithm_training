package sort

func choosePivotIndex(arr []int, begin, end int) int {
	return begin
}

func startQuicksort(arr []int, begin, end int) {
	if begin >= end {
		return
	}
	curBegin, curEnd := begin, end
	for curBegin <= curEnd {
		for arr[curBegin] < arr[begin] {
			curBegin++
		}
		for arr[curEnd] > arr[begin] {
			curEnd--
		}
		if curBegin <= curEnd {
			arr[curBegin], arr[curEnd] = arr[curEnd], arr[curBegin]
			curBegin++
			curEnd--
		}
	}
	startQuicksort(arr, begin, curEnd)
	startQuicksort(arr, curBegin, end)
}

func QuickSort(arr []int) []int {
	copyarray := make([]int, len(arr))
	copy(copyarray, arr)
	startQuicksort(copyarray, 0, len(arr)-1)
	return copyarray
}

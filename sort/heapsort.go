package sort

func HeapSort(arr []int) []int {
	copyArray := make([]int, len(arr))
	copy(copyArray, arr)
	startHeapSort(copyArray)
	return copyArray
}

func startHeapSort(arr []int) {
	for i := (len(arr) - 2) / 2; i >= 0; i-- {
		downHeap(arr, i, len(arr)-1)
	}
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		downHeap(arr, 0, i-1)
	}
}

func downHeap(arr []int, cursolParent, end int) {
	var (
		cursolChild int
		parent      = arr[cursolParent]
	)

	for {
		cursolChild = 2*cursolParent + 1
		if cursolChild > end {
			break
		}
		if cursolChild != end {
			if arr[cursolChild+1] > arr[cursolChild] {
				cursolChild = cursolChild + 1
			}
		}
		if parent >= arr[cursolChild] {
			break
		}
		arr[cursolParent] = arr[cursolChild]
		cursolParent = cursolChild
	}
	arr[cursolParent] = parent
}

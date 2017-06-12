package sort

func SelectSort(arr []int) []int {
	copyarry := make([]int, len(arr))
	copy(copyarry, arr)

	var min int
	for i := 0; i < len(copyarry)-1; i++ {
		min = i
		for j := i + 1; j < len(copyarry); j++ {
			if copyarry[j] < copyarry[min] {
				min = j
			}
		}
		copyarry[i], copyarry[min] = copyarry[min], copyarry[i]
	}
	return copyarry
}

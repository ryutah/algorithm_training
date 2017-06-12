package sort

func InsertSort(arr []int) []int {
	copyarray := make([]int, len(arr))
	copy(copyarray, arr)

	for i := 1; i < len(copyarray); i++ {
		for j := i; j > 0 && copyarray[j-1] > copyarray[j]; j-- {
			copyarray[j-1], copyarray[j] = copyarray[j], copyarray[j-1]
		}
	}
	return copyarray
}

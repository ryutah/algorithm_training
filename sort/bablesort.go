package sort

func BableSort(arr []int) []int {
	copyarray := make([]int, len(arr))
	copy(copyarray, arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := len(arr) - 1; j > i; j-- {
			if copyarray[i] > copyarray[j] {
				copyarray[i], copyarray[j] = copyarray[j], copyarray[i]
			}
		}
	}
	return copyarray
}

package sorting

func BubbleSort(arr []int) []int {

	n := len(arr)

	for i := 0; i < n-1; i++ {

		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				// swap arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

	}
	return arr
}

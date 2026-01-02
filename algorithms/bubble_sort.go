package algorithms

func BubbleSort(arr []int) {

	size := len(arr)
	limit := size

	for limit > 0 {
		for i := 0; i < limit-1; i++ {

			if arr[i] > arr[i+1] {
				tmp := arr[i+1]
				arr[i+1] = arr[i]
				arr[i] = tmp
			}
		}

		limit--
	}
}

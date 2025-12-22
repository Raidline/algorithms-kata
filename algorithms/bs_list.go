package algorithms

func BinarySearch(haystack []int, needle int) bool {

	lo := 0
	hi := len(haystack)

	for lo < hi {
		mid := (lo + (hi-lo)/2)
		value := haystack[mid]

		if value == needle {
			return true
		}

		if value < needle {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return false
}

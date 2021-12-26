package binarysearch

func SearchInts(list []int, key int) int {

	lo, hi := 0, len(list) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2

		if list[mid] == key {
			return mid
		} else if list[mid] < key {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return -1
}

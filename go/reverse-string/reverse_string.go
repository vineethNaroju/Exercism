package reverse

func Reverse(input string) string {
	rin := []rune(input)

	for lo, hi := 0, len(rin)-1; lo < hi; lo, hi = lo+1, hi-1 {
		rin[lo], rin[hi] = rin[hi], rin[lo]
	}

	return string(rin)
}

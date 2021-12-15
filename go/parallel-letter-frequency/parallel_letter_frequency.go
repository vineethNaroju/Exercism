package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	mp := FreqMap{}

	mpChannel := make(chan FreqMap, len(l))

	for _, str := range l {
		go func(s string) {
			mpChannel <- Frequency(s)
		}(str)
	}

	for range l {
		m := <-mpChannel
		for k, v := range m {
			mp[k] += v
		}
	}

	return mp
}

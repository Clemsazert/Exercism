package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := make(FreqMap, 48)
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	channel := make(chan FreqMap, len(l))
	for i := 0; i < len(l); i++ {
		i := i
		go func() {
			channel <- Frequency(l[i])
		}()
	}
	result := make(FreqMap, 48)
	for i := 0; i < len(l); i++ {
		for key, amount := range <-channel {
			result[key] += amount
		}
	}
	return result
}

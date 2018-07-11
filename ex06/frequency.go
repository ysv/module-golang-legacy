package letter

import "sync"

func Frequency(text string) map[rune]int {
	m := map[rune]int{}

	for _, v := range text {
		m[v] += 1
	}

	return m
}

type SafeMap struct {
	buf map[rune]int
	*sync.Mutex
}

func concFrequency(text string, fmap *SafeMap, done chan byte) {
	for _, v := range text {
		fmap.Lock()
		fmap.buf[v] += 1
		fmap.Unlock()
	}
	done <- 1
}

func ConcurrentFrequency(texts []string) map[rune]int {
	fmap := &SafeMap{map[rune]int{}, &sync.Mutex{}}
	len := len(texts)
	done := make(chan byte, len)

	for _, text := range texts {
		go concFrequency(text, fmap, done)
	}

	for i := 0; i < len; i++ {
		<-done
	}

	return fmap.buf
}

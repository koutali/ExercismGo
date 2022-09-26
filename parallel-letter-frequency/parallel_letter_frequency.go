package letter

import (
	"fmt"
	"sync"
)

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

var mutex = &sync.RWMutex{}

func ParallelFrequency(s string, occurences FreqMap, channel chan string) {

	mutex.Lock()
	for _, runic := range s {
		occurences[runic]++
	}
	mutex.Unlock()
	channel <- "Done"
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(text []string) FreqMap {
	channel := make(chan string)
	occurences := FreqMap{}
	for _, subText := range text {
		go ParallelFrequency(subText, occurences, channel)
	}

	for i := 0; i < len(text); i++ {
		fmt.Println(<-channel)
	}

	return occurences
}

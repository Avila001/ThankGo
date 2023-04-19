package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

// counterGo1 stores the number of digits in each word.
// each key is a word and value is the number of digits in the word.
type counterGo1 map[string]int

// countDigitsInWordsGo1 counts digits in pharse words
func countDigitsInWordsGo1(phrase string) counterGo1 {
	words := strings.Fields(phrase)
	syncStats := sync.Map{}

	var wg sync.WaitGroup

	// начало решения
	//counts := make([]int, 4)
	// Посчитайте количество цифр в словах,
	// используя отдельную горутину для каждого слова.
	wg.Add(len(words))
	for _, word := range words {
		go func(word string) {
			defer wg.Done()
			count := countDigitsGo1(word)
			syncStats.Store(word, count)
		}(word)
	}
	// Чтобы записать результаты подсчета,
	// используйте syncStats.Store(word, count)

	// В результате syncStats должна содержать слова
	// и количество цифр в каждом.
	wg.Wait()
	// конец решения

	return asStats(syncStats)
}

// countDigitsGo1 returns the number of digits in a string
func countDigitsGo1(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// asStats converts stats from sync.Map to ordinary map
func asStats(m sync.Map) counterGo1 {
	stats := counterGo1{}
	m.Range(func(word, count any) bool {
		stats[word.(string)] = count.(int)
		return true
	})
	return stats
}

// printStatsGo1 prints words and their digit counts
func printStatsGo1(stats counterGo1) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func mainGoroutine1() {
	phrase := "0ne 1wo thr33 4068"
	counts := countDigitsInWordsGo1(phrase)
	printStatsGo1(counts)
}

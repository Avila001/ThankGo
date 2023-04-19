package main

import (
	"fmt"
	"strings"
	"unicode"
)

// counterGo1 хранит количество цифр в каждом слове.
// ключ карты - слово, а значение - количество цифр в слове.
type counterGo2 map[string]int

// countDigitsInWordsGo1 считает количество цифр в словах фразы
func countDigitsInWordsGo2(phrase string) counterGo2 {
	words := strings.Fields(phrase)
	counted := make(chan int)

	// начало решения
	stats := counterGo2{}
	go func() {
		// Пройдите по словам,
		// посчитайте количество цифр в каждом,
		// и запишите его в канал counted
		for _, word := range words {
			count := countDigitsGo2(word)
			counted <- count
		}

	}()
	// Считайте значения из канала counted
	// и заполните stats.
	for _, word := range words {
		stats[word] = <-counted
	}
	// В результате stats должна содержать слова
	// и количество цифр в каждом.

	// конец решения

	return stats
}

// countDigitsGo1 возвращает количество цифр в строке
func countDigitsGo2(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// printStatsGo1 печатает слова и количество цифр в каждом
func printStatsGo2(stats counterGo2) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func mainGo2() {
	phrase := "0ne 1wo thr33 4068"
	stats := countDigitsInWordsGo2(phrase)
	printStatsGo2(stats)
}

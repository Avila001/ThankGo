package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// element - интерфейс элемента последовательности
type element interface{}

// weightFunc - функция, которая возвращает вес элемента
type weightFunc func(element) int

// iterator - интерфейс, который умеет
// поэлементно перебирать последовательность
type iterator interface {
	// чтобы понять сигнатуры методов - посмотрите,
	// как они используются в функции max() ниже
	next() bool
	val() element
}

// intIterator - итератор по целым числам
// (реализует интерфейс iterator)
type intIterator struct {
	// поля структуры
	index int
	items []int
}

// методы intIterator, которые реализуют интерфейс iterator
func (it *intIterator) next() bool {
	if it.index < len(it.items) {
		return true
	}
	return false
}

func (it *intIterator) val() element {
	item := it.items[it.index]
	it.index++
	return item
}

// конструктор intIterator
func newIntIterator(src []int) *intIterator {
	// ...
	return &intIterator{items: src}
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// main находит максимальное число из переданных на вход программы.
//func main() {
//	//nums := readInputIterator()
//	nums := []int{1, 2, 3, 4, 5}
//	it := newIntIterator(nums)
//	weight := func(el element) int {
//		return el.(int)
//	}
//	m := max(it, weight)
//	fmt.Println(m)
//}

// max возвращает максимальный элемент в последовательности.
// Для сравнения элементов используется вес, который возвращает
// функция weight.
func max(it iterator, weight weightFunc) element {
	var maxEl element
	for it.next() {
		//fmt.Println(it.val())
		curr := it.val()
		if maxEl == nil || weight(curr) > weight(maxEl) {
			maxEl = curr
		}
	}
	return maxEl
}

// readInputErrors считывает последовательность целых чисел из os.Stdin.
func readInputIterator() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	return nums
}

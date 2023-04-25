package main

/*
import (
	"fmt"
	"strconv"
	"testing"
	"unicode"
)

// начало решения

// calcDistance возвращает общую длину маршрута в метрах
func calcDistance(directions []string) int {
	var a int
	for _, str := range directions {
		a += int(calsDistanceFromString(str))
	}
	return a
}

func calsDistanceFromString(str string) float64 {
	//var distance int
	var distString string
	var index int
	var res float64
	for i := 0; i < len(str); i++ {
		if unicode.IsDigit(rune(str[i])) || rune(str[i]) == '.' {
			distString += string(rune(str[i]))
			if unicode.IsLetter(rune(str[i+1])) {
				index = i + 1
			}
		}
	}
	//fmt.Println("index", string(str[index]))
	if str[index] == 'm' {
		res, _ = strconv.ParseFloat(distString, 64)
	}

	if str[index] == 'k' {
		res, _ = strconv.ParseFloat(distString, 64)
		fmt.Println(res)
		res = res * 1000
		fmt.Println(res)
	}
	return res
}

// конец решения

func main() {
	directions := []string{
		"straight 1.6km",
	}
	const want = 16000
	got := calcDistance(directions)
	fmt.Println(got)
	if got != want {
		fmt.Printf("%v: got %v, want %v", directions, got, want)
	}
}

func Test(t *testing.T) {
	directions := []string{
		"100m to intersection",
		"turn right",
		"straight 300m",
		"enter motorway",
		"straight 5km",
		"exit motorway",
		"500m straight",
		"turn sharp left",
		"continue 100m to destination",
	}
	fmt.Printf("%T", directions[0])
	//const want = 6000
	//got := calcDistance(directions)
	//if got != want {
	//	t.Errorf("%v: got %v, want %v", directions, got, want)
	//}
}
*/

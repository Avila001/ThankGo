package main

// не удаляйте импорты, они используются при проверке
import (
	"testing"
)

type Words struct {
	str string
}

func MakeWords(s string) Words {
	return Words{s}
}

func (w Words) Index(word string) int {
	str := w.str
	count := 0
	for i := 0; i < len(str); i++ {
		// Если текущий символ - пробел или знак пунктуации, значит, мы нашли новое слово
		if str[i] == ' ' || str[i] == ',' || str[i] == '!' || str[i] == '?' || str[i] == '.' {
			count++
		} else {
			// Если текущий символ не является пробелом или знаком пунктуации,
			// значит, он является частью слова, поэтому мы сравниваем его с искомым словом
			if i+len(word) <= len(str) && str[i:i+len(word)] == word {
				return count
				break
			}
		}
	}
	return -1
}

func BenchmarkIndex(b *testing.B) {
	w := Words{"a"}
	mk := MakeWords("in a coat of gold or a coat of red")
	for i := 0; i < b.N; i++ {
		w.Index(mk.str)
	}
}

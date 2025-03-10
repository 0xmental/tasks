package mapTasks

import (
	"strings"
	"unicode"
)

/*
Напиши функцию, которая принимает текст (строку) и возвращает мапу map[string]int,
где ключ — это слово, а значение — частота его появления. Игнорируй регистр и знаки препинания.
*/

func WordFreq(inputText string) map[string]int {
	inputText = strings.ToLower(inputText)
	var word []rune
	freq := make(map[string]int)

	for _, v := range inputText {
		if unicode.IsLetter(v) {
			word = append(word, v)
		} else if v == '-' {
			if len(word) > 0 {
				word = append(word, v)
			}
		} else {
			if len(word) > 0 {
				freq[string(word)]++
				word = nil
			}
		}
	}

	return freq
}

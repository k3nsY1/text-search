package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"unicode"
)

func main() {

	file, err := os.Open("mobydick.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var words [][]rune
	var word []rune
	var endword bool
	s := bufio.NewScanner(file)
	for s.Scan() {
		for _, r := range bytes.Runes(s.Bytes()) {
			if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' && !endword {
				word = append(word, unicode.ToLower(r))
			}
			if unicode.IsSpace(r) {
				endword = true
			}
			if endword {
				words = append(words, word)
				word = nil
				endword = false
			}
		}

	}

	var first [][]rune
	var second []int
	// // Пробегаемся по циклу и если words не нулевой, то сравниваем два элемента, первый со следующим
	for i := 0; i < len(words); i++ {
		if words[i] != nil {
			var cnt = 1
			for j := i + 1; j < len(words); j++ {

				res := compare(words[i], words[j]) //Compare(words[i], words[j])
				// Если равны, то добавляем в счетчик и удаляем элемент
				if res == true {
					cnt++
					words[j] = nil
				}
			}
			// В первый добавляем слово
			first = append(first, words[i])
			// Во второй количество
			second = append(second, cnt)
		}
	}
	// // Сортируем массивы по количеству
	for i := 0; i < len(first); i++ {
		for j := i + 1; j < len(first); j++ {
			if second[i] < second[j] {
				second[i], second[j] = second[j], second[i]
				first[i], first[j] = first[j], first[i]
			}
		}
	}
	// // Выводим 20 элементов
	for i := 0; i < len(first); i++ {
		fmt.Println(first[i], string(first[i]), second[i])
		if i == 19 {
			break
		}
	}

}

func compare(a []rune, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for _, v := range a {
		for _, c := range b {
			if v != c {
				return false
			}
		}
	}
	return true
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	// Открываем файл для чтения
	file, err := os.Open("mobydick.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var words [][]byte
	var word []byte
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		// Запись слов в нижний регистр и избавление от знаков
		if word != nil {
			words = append(words, bytes.ToLower(word))
			word = nil
		}
		// fmt.Println(scanner.Text())
		for i := 0; i < len(scanner.Bytes()); i++ {
			if scanner.Bytes()[i] >= 'a' && scanner.Bytes()[i] <= 'z' || scanner.Bytes()[i] >= 'A' && scanner.Bytes()[i] <= 'Z' {
				word = append(word, scanner.Bytes()[i])
			}
		}
	}
	// for _, v := range words {
	// 	fmt.Println(string(v))
	// }

	var uniq [][]byte
	var count []int
	// Пробегаемся по циклу и если words не нулевой, то сравниваем два элемента, первый со следующим
	for i := 0; i < len(words); i++ {
		if words[i] != nil {
			var cnt = 1
			for j := i + 1; j < len(words); j++ {
				res := bytes.Compare(words[i], words[j])
				// Если равны, то добавляем в счетчик и удаляем элемент
				if res == 0 {
					cnt++
					words[j] = nil
				}
			}
			// В первый добавляем слово
			uniq = append(uniq, words[i])
			// Во второй количество
			count = append(count, cnt)
		}
	}
	// Сортируем массивы по количеству
	for i := 0; i < len(uniq); i++ {
		for j := i + 1; j < len(uniq); j++ {
			if count[i] < count[j] {
				count[i], count[j] = count[j], count[i]
				uniq[i], uniq[j] = uniq[j], uniq[i]
			}
		}
	}
	// Выводим 20 элементов
	for i := 0; i < len(uniq); i++ {
		fmt.Println(string(uniq[i]), count[i])
		if i == 19 {
			break
		}
	}
}

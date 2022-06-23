package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// type Word struct {
// 	Bytes []byte
// 	// word []byte
// 	Count int
// }

func main() {
	file, err := os.Open("mobydick.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var words [][]byte
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		// words := scanner.Text()

		// Запись слов в нижний регистр и избавление от знаков
		for i := 0; i < len(scanner.Bytes()); i++ {
			if scanner.Bytes()[i] >= 65 && scanner.Bytes()[i] <= 90 {
				scanner.Bytes()[i] = scanner.Bytes()[i] + 32
				words = append(words, scanner.Bytes())
			} else if scanner.Bytes()[i] == 32 || scanner.Bytes()[i] == 42 {
				scanner.Bytes()[i] = 0
				words = append(words, scanner.Bytes())
			} else {
				words = append(words, scanner.Bytes())
			}
		}
		// fmt.Println(scanner.Text())
	}
	// for i := 0; i < len(words); i++ {
	// 	fmt.Println(words[i])
	// }

	var first [][]byte
	var second []int
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
			first = append(first, words[i])
			// Во второй количество
			second = append(second, cnt)
		}
	}
	// Сортируем массивы по количеству
	for i := 0; i < len(first); i++ {
		for j := i + 1; j < len(first); j++ {
			if second[i] < second[j] {
				second[i], second[j] = second[j], second[i]
				first[i], first[j] = first[j], first[i]
			}
		}
	}
	// Выводим 20 элементов
	for i := 0; i < len(first); i++ {
		fmt.Println(string(first[i]), second[i])
		if i == 19 {
			break
		}
	}

	// resArray := []Word{}
	// for i := 0; i < len(words); i++ {
	// 	currentWord := Word{
	// 		Bytes: words[i],
	// 		Count: 1,
	// 	}
	// 	for j := i + 1; j < len(words); j++ {
	// 		res := bytes.Compare(words[i], words[j])
	// 		if res == 0 {
	// 			currentWord.Count++
	// 			words[j] = nil
	// 		}
	// 	}
	// 	resArray = append(resArray, currentWord)
	// 	fmt.Println(string(resArray[i].Bytes), " : ", resArray[i].Count)
	// }
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type Word struct {
	Bytes []byte
	// word []byte
	Count int
}

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
		// r, _ := utf8.DecodeRune(words[])
		words = append(words, scanner.Bytes())

		// fmt.Println(scanner.Text())
	}
	// for i := 0; i < len(words); i++ {
	// 	fmt.Println(words[i])
	// }
	var first [][]byte
	var second []int
	for i := 0; i < len(words); i++ {
		if words[i] != nil {
			var cnt = 1
			for j := i + 1; j < len(words); j++ {
				res := bytes.Compare(words[i], words[j])
				if res == 0 {
					cnt++
					words[j] = nil
				}
			}
			first = append(first, words[i])
			second = append(second, cnt)
		}
	}
	for i := 0; i < len(first); i++ {
		for j := i + 1; j < len(first); j++ {
			if second[i] < second[j] {
				second[i], second[j] = second[j], second[i]
				first[i], first[j] = first[j], first[i]
				// swa(second[i], second[j])
			}
		}
	}
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

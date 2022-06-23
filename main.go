package main

import (
	"bufio"
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

		fmt.Println(scanner.Text())
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

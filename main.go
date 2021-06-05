package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	dict map[string]int
)

func main() {
	filePath := "dict.txt"
	readFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Failed to read File")
	}
	f := bufio.NewScanner(readFile)

	f.Split(bufio.ScanWords)
	dict = make(map[string]int)
	for f.Scan() {
		dict[f.Text()] = 1
	}
	sentence := "backdoordoor"
	for _, v := range wordBreak(sentence) {
		fmt.Println(v)
	}

}

func wordBreak(sentence string) []string {

	var result []string
	if dict[sentence] == 1 {
		result = append(result, sentence)
	}
	for i := 1; i < len(sentence); i++ {
		prefix := sentence[0:i]
		if dict[prefix] == 1 {
			returnStringsList := wordBreak(sentence[i:])
			for _, v := range returnStringsList {
				result = append(result, prefix+" "+v)
			}
		}
	}
	return result
}

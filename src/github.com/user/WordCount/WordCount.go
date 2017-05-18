package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)

func WordCount(s string) map[string]int {
	var m map[string]int = make(map[string]int)
	keys := strings.Fields(s)
	for i,key := range keys { 
		_ = i
		word, ok := m[key]
		_ = word
		if (ok) { 
			m[key]++
		} else { 
			m[key] = 1
		}
	}
	return m
}

func main() {
	// read the input
	fmt.Println("Please input a sentence: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadBytes('\n') // read until enter

	// convert input to string
	var sentence string
	sentence = string(input[0:len(input)-1]) // cut off the '\n'
	result := WordCount(sentence) // call the function
	fmt.Println("The result is: \n", result)
}

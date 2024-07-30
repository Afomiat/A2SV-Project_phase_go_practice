package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func wordFreq(s string) map[string]int {
	
	s = strings.ToLower(s)

	re := regexp.MustCompile(`[^\w\s]`)
	s = re.ReplaceAllString(s, "")

	words := strings.Fields(s)

	wordCount := make(map[string]int)

	for _, word := range words {
		
		wordCount[word]++
	}

	return wordCount
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str) 
	
	frequencies := wordFreq(str)

	fmt.Println("Word frequencies:")
	for word, count := range frequencies {
		fmt.Printf("%s: %d\n", word, count) 
	}
}

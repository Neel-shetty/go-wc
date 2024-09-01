package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func countWords(line string, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	words := strings.Fields(line)
	c <- len(words)
}

func main() {
	fmt.Println("Enter text (Ctrl-D to finish typing):")
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	wordsChan := make(chan int)
	lineCount := 0
	var wg sync.WaitGroup

	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			wg.Add(1)
			go countWords(line, wordsChan, &wg)
			input += line + "\n"
			lineCount++
		}
		wg.Wait()
		close(wordsChan)
	}()

	totalWords := 0
	for wordCount := range wordsChan {
		totalWords += wordCount
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input = strings.TrimSpace(input)
	charCount := len(input) - strings.Count(input, "\n")

	fmt.Printf("\nTotal characters: %d\n", charCount)
	fmt.Printf("Total words: %d\n", totalWords)
	fmt.Printf("Total lines: %d\n", lineCount)
}


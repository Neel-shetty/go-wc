package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter a string")
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		input += line + "\n"
	}
	fmt.Println("The input was:", input)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	var numChars int
	var numWords int
	var numLines int
	for i, c := range input {
		fmt.Println(i, c)
		if c == 10 {
			numWords += 1
			numLines += 1
			continue
		}
		if c == 32 {
			numWords += 1
			continue
		}
		numChars += 1
	}
	fmt.Println("total chars:", numChars)
	fmt.Println("total words:", numWords)
	fmt.Println("total lines:", numLines)
}

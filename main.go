package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open dict file
	file, err := os.Open("/var/lib/words/dict")
	if err != nil {
		fmt.Println("Error: could not open file:", err)
		return
	}
	defer file.Close()

	// Create array to count letters (26 slots for a-z)
	var letterCounts [26]int

	// Use scanner to read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // Get current line as string
		for _, char := range line {
			// Check if character is a letter and update counts
			if char >= 'a' && char <= 'z' {
				letterCounts[char-'a']++
			} else if char >= 'A' && char <= 'Z' {
				letterCounts[char-'A']++
			}
		}
	}

	// Check for errors during file scan
	if err := scanner.Err(); err != nil {
		fmt.Println("Error while reading file:", err)
		return
	}

	// Print results (a-z)
	fmt.Println("Frequency of each letter:")
	for i := 0; i < len(letterCounts); i++ {
		fmt.Printf("%c: %d\n", 'a'+i, letterCounts[i])
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	input_file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error while opening input file")
	}

	output_file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal("error while creating file")
	}

	scanner := bufio.NewScanner(input_file)
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(output_file)

	for scanner.Scan() {
		word := scanner.Text()
		if isPalindrome(word) {
			fmt.Fprintln(writer, word)
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error during scan: ", err)
	}

	writer.Flush()

}

func isPalindrome(word string) bool {

	runes := []rune(word)
	for i, j := 0, len(word)-1; i < j; {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}

	rev := string(runes)
	if word == rev {
		return true
	}
	return false

}

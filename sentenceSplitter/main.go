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
		log.Fatal("error while reading")
	}
	defer input_file.Close()

	output_file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal("error while creating file")
	}
	defer output_file.Close()

	scanner := bufio.NewScanner(input_file)
	scanner.Split(bufio.ScanRunes)
	writer := bufio.NewWriter(output_file)

	isNewLine := false

	for scanner.Scan() {
		char := scanner.Text()
		if isNewLine {
			fmt.Fprint(writer, "\n")
			isNewLine = false
			continue
		}
		fmt.Fprint(writer, char)
		if char == "." {
			isNewLine = true
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal("error while scan")
	}

	err = writer.Flush()
	if err != nil {
		log.Fatal("error while flush")
	}

}

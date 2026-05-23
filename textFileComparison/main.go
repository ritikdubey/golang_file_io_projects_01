package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	input_file1, err1 := os.Open("input1.txt")
	input_file2, err2 := os.Open("input2.txt")

	if err1 != nil {
		fmt.Println("error opening files: ", err2)
	}

	if err2 != nil {
		fmt.Println("error opening files: ", err2)
	}

	output_file, err3 := os.Create("output_file.txt")

	if err3 != nil {
		fmt.Println("error while creating file: ", err3)
	}

	scanner1 := bufio.NewScanner(input_file1)
	scanner2 := bufio.NewScanner(input_file2)

	writer := bufio.NewWriter(output_file)

	slice1 := []string{}
	slice2 := []string{}

	for scanner1.Scan() {
		line1 := scanner1.Text()
		slice1 = append(slice1, line1)
	}

	for scanner2.Scan() {
		line2 := scanner2.Text()
		slice2 = append(slice2, line2)
	}

	arrSize := len(slice2)
	if len(slice1) > len(slice2) {
		arrSize = len(slice1)
	}

	for i := 0; i < arrSize; i++ {

		var line1, line2 string

		if i < len(slice1) {
			line1 = slice1[i]
		}

		if i < len(slice2) {
			line2 = slice2[i]
		}

		if i > 0 {
			fmt.Fprintf(writer, "\n")
		}

		if strings.EqualFold(line1, line2) {
			fmt.Fprintf(writer, "No difference found: %v", line1)
		} else {
			fmt.Fprintf(writer, "Difference found, line1: %v, line2: %v", line1, line2)
		}
	}

	err4 := writer.Flush()
	if err4 != nil {
		fmt.Println("error while writing file: ", err4)
	}

}

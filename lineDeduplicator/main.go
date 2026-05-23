package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {

	caseSensitivityEnabled := true
	trimmingEnabled := true

	input_file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("error while reading the file: ", err)
		return
	}

	output_file, err := os.Create("output.txt")

	if err != nil {
		fmt.Println("error while creating output file: ", err)
	}

	scanner := bufio.NewScanner(input_file)
	writer := bufio.NewWriter(output_file)

	linesArr := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if trimmingEnabled {
			line = strings.TrimSpace(line)
		}
		if caseSensitivityEnabled {
			line = strings.ToLower(line)
		}

		if slices.Contains(linesArr, line) || line == "" {
			continue
		} else {
			linesArr = append(linesArr, line)
			fmt.Fprintf(writer, "%v\n", line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error while scan: ", err)
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("error while flush: ", err)
	}

}

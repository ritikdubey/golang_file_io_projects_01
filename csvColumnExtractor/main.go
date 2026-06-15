package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	input_file, err := os.Open("input.csv")
	if err != nil {
		log.Fatalf("error while opening file: %v", err)
	}

	output_file, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error while creating file: %v", err)
	}

	reader := csv.NewReader(input_file)
	writer := bufio.NewWriter(output_file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("error while reading data: %v", err)
	}

	columnToSearch := "City"
	index := -1

	for i := 0; i < len(records[0]); i++ {
		if records[0][i] == columnToSearch {
			index = i
		}
	}

	if index == -1 {
		log.Fatal("selected column does not exist")
	}

	for i := 1; i < len(records); i++ {
		fmt.Fprintf(writer, "%v\n", records[i][index])
	}

	if err = writer.Flush(); err != nil {
		log.Fatalf("error while flush: %v", err)
	}

}

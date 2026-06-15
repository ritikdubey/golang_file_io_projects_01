package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func main() {

	input1, err := os.Open("input1.csv")
	if err != nil {
		log.Fatalf("error while reading file: %v", err)
	}

	input2, err := os.Open("input2.csv")
	if err != nil {
		log.Fatalf("error while reading file: %v", err)
	}

	output, err := os.Create("output.csv")

	reader1 := csv.NewReader(input1)
	reader2 := csv.NewReader(input2)

	writer := csv.NewWriter(output)

	data1, err := reader1.ReadAll()
	if err != nil {
		log.Fatalf("error while reading data: %v", err)
	}

	data2, err := reader2.ReadAll()
	if err != nil {
		log.Fatalf("error while reading data: %v", err)
	}

	if strings.Join(data1[0], "") != strings.Join(data2[0], "") {
		log.Fatal("column name not same")
		return
	}

	newRecords := append(data1, data2[1:]...)

	if err = writer.WriteAll(newRecords); err != nil {
		log.Fatalf("error while writing data to file: %v", err)
	}

}

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	input_file, err := os.Open("data.csv")
	if err != nil {
		log.Fatalf("error while opening file: %v", err)
	}
	defer input_file.Close()

	output_file, err := os.Create("report.txt")
	if err != nil {
		log.Fatalf("error while creating output file: %v", err)
	}
	defer output_file.Close()

	reader := csv.NewReader(input_file)
	writer := bufio.NewWriter(output_file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("error while reading records: %v", err)
	}

	for rowIndex, row := range records[1:] {

		if rowIndex == 0 {
			fmt.Fprintf(writer, "------------------------------------------------\n")
		}

		fmt.Fprintf(writer, "Name: %v\n", row[1])
		fmt.Fprintf(writer, "Maths: %v\n", row[2])
		fmt.Fprintf(writer, "Science: %v\n", row[3])
		fmt.Fprintf(writer, "English: %v\n", row[4])
		fmt.Fprintf(writer, "History: %v\n", row[5])
		fmt.Fprintf(writer, "Computer: %v\n", row[6])

		math, err := strconv.Atoi(row[2])
		if err != nil {
			log.Fatalf("error while conversion: %v", err)
		}

		science, err := strconv.Atoi(row[3])
		if err != nil {
			log.Fatalf("error while conversion: %v", err)
		}

		english, err := strconv.Atoi(row[3])
		if err != nil {
			log.Fatalf("error while conversion: %v", err)
		}

		history, err := strconv.Atoi(row[3])
		if err != nil {
			log.Fatalf("error while conversion: %v", err)
		}

		computer, err := strconv.Atoi(row[3])
		if err != nil {
			log.Fatalf("error while conversion: %v", err)
		}

		sum := math + science + english + history + computer
		avg := (float64(sum) / 500) * 100

		fmt.Fprintf(writer, "\n")
		fmt.Fprintf(writer, "Percentage: %.2f\n", avg)

		fmt.Fprintf(writer, "------------------------------------------------\n")
	}

	if err = writer.Flush(); err != nil {
		log.Fatalf("error while writing file: %v", err)
	}

}

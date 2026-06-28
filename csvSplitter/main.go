package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func main() {

	input_file, err := os.Open("input.csv")
	if err != nil {
		log.Fatalf("error while reading file: %v", err)
	}
	defer input_file.Close()

	reader := csv.NewReader(input_file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("error while reading records: %v", err)
	}

	header := records[0]

	counter := 0
	file_count := 1

	output_file, err := os.Create("output_" + strconv.Itoa(file_count) + ".csv")
	if err != nil {
		log.Fatalf("error while creating file: %v", err)
	}

	writer := csv.NewWriter(output_file)
	writer.Write(header)

	for _, row := range records[1:] {

		if counter == 0 && file_count != 1 {

			writer.Flush()
			if err = writer.Error(); err != nil {
				log.Fatalf("error while flush: %v", err)
			}

			output_file.Close()

			output_file, err = os.Create("output_" + strconv.Itoa(file_count) + ".csv")
			if err != nil {
				log.Fatalf("error while creating file: %v", err)
			}
			writer = csv.NewWriter(output_file)
			writer.Write(header)
		}

		counter = counter + 1

		if counter == 10 {
			counter = 0
			file_count = file_count + 1
		}

		err = writer.Write(row)
		if err != nil {
			log.Fatalf("error while updating rows: %v", err)
		}
	}

	writer.Flush()
	if err = writer.Error(); err != nil {
		log.Fatalf("error while flush: %v", err)
	}

	output_file.Close()

}

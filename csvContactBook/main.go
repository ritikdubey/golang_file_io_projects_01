package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	csv_file, err := os.OpenFile("contactbook.csv", os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("error while reading file: %s", err)
	}
	defer csv_file.Close()

	reader := csv.NewReader(csv_file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("failed to parse csv: %s", err)
	}

	fmt.Println(records)

	writer := csv.NewWriter(csv_file)

	var serialNum string
	var firstName string
	var lastName string
	var emailId string
	var contactNum string
	var city string

	fmt.Println("Enter serialNum, first name, last name, email id, contact num, city (separated by space): ")
	_, err = fmt.Scan(&serialNum, &firstName, &lastName, &emailId, &contactNum, &city)

	newRecords := [][]string{{serialNum, firstName, lastName, emailId, contactNum, city}}

	// csvData := [][]string{}
	// csvData = append(csvData, newRecords...)

	if err = writer.WriteAll(newRecords); err != nil {
		log.Fatal("Error writing all records:", err)
	}

	var firstNameToSearch string = "Dan"

	for _, row := range records {
		for columnIndex, column := range row {
			//fmt.Println(rowIndex, columnIndex, column)
			if columnIndex == 1 {
				if strings.Contains(column, firstNameToSearch) {
					fmt.Println(row)
				}
			}
		}
	}

}

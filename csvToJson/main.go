package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type User struct {
	Id          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Department  string `json:"department"`
	Salary      int    `json:"salary"`
	JoiningDate string `json:"joiningDate"`
	City        string `json:"city"`
	Active      bool   `json:"active"`
}

func main() {

	input_file, err := os.Open("data.csv")
	if err != nil {
		log.Fatalf("Error while reading file: %v", err)
	}
	defer input_file.Close()

	outfile_file, err := os.Create("output.json")
	if err != nil {
		log.Fatalf("error while creating file: %v", err)
	}
	defer outfile_file.Close()

	reader := csv.NewReader(input_file)
	encoder := json.NewEncoder(outfile_file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error while parsing data: %v", err)
	}

	dataArr := []User{}

	for rowIndex, row := range records {
		if rowIndex == 0 {
			continue
		}
		id, err := strconv.Atoi(row[0])
		if err != nil {
			log.Fatal("invalid ID")
		}
		salary, err := strconv.Atoi(row[5])
		if err != nil {
			log.Fatal("invalid salary")
		}
		active, err := strconv.ParseBool(row[8])
		if err != nil {
			log.Fatal("invalid active status")
		}

		user := User{id, row[1], row[2], row[3], row[4], salary, row[6], row[7], active}

		dataArr = append(dataArr, user)

	}

	fmt.Println(dataArr)

	if err = encoder.Encode(dataArr); err != nil {
		log.Fatalf("error while encoding: %v", err)
	}

}

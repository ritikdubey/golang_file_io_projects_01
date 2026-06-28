package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"net/mail"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {

	data_file, err := os.Open("data.csv")
	if err != nil {
		log.Fatalf("error while reading file: %v", err)
	}
	defer data_file.Close()

	output_file, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error while creating file: %v", err)
	}
	defer output_file.Close()

	reader := csv.NewReader(data_file)
	writer := bufio.NewWriter(output_file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("error while reading records: %v", err)
	}

	for ri, row := range records {
		if ri == 0 {
			continue
		}
		if len(strings.TrimSpace(row[0])) == 0 || !strings.HasPrefix(row[0], "EMP") {
			fmt.Fprintf(writer, "Row %v: Employee Id is not not valid\n", ri)
		}
		if len(strings.TrimSpace(row[1])) == 0 || strings.ContainsFunc(row[1], unicode.IsDigit) {
			fmt.Fprintf(writer, "Row %v: Name is not valid\n", ri)
		}

		if len(strings.TrimSpace(row[2])) == 0 || !isEmailValid(row[2]) {
			fmt.Fprintf(writer, "Row %v: Email is not valid\n", ri)
		}
		if len(strings.TrimSpace(row[3])) == 0 || strings.ContainsFunc(row[3], unicode.IsLetter) || len(row[3]) != 10 {
			fmt.Fprintf(writer, "Row %v: Phone no. is not valid\n", ri)
		}
		age, _ := strconv.Atoi(row[4])
		if len(strings.TrimSpace(row[4])) == 0 || age < 18 || age > 65 {
			fmt.Fprintf(writer, "Row %v: Age is not valid\n", ri)
		}
		if len(strings.TrimSpace(row[5])) == 0 || !isDepartmentValid(row[5]) {
			fmt.Fprintf(writer, "Row %v: Department is not valid\n", ri)
		}
		salary, _ := strconv.Atoi(row[6])
		if len(strings.TrimSpace(row[6])) == 0 || salary < 0 {
			fmt.Fprintf(writer, "Row %v: Salary is not valid\n", ri)
		}
		// 01 (Month) / 02 (Day) / 03 (Hour) / 04 (Minute) / 05 (Second) / 06 (Year).
		layout1 := "2006-01-02"
		layout2 := "2006-02-01"
		_, err1 := time.Parse(layout1, row[7])
		_, err2 := time.Parse(layout2, row[7])
		if len(strings.TrimSpace(row[7])) == 0 || (err1 != nil && err2 != nil) {
			fmt.Fprintf(writer, "Row %v: JoinDate is not valid\n", ri)
		}
		if len(strings.TrimSpace(row[8])) == 0 || (row[8] != "Active" && row[8] != "Inactive") {
			fmt.Fprintf(writer, "Row %v: Status is not valid\n", ri)
		}
	}

	if err := writer.Flush(); err != nil {
		log.Fatalf("error while writing file: %v", err)
	}

}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isDepartmentValid(department string) bool {
	if department == "Engineering" || department == "HR" || department == "Finance" ||
		department == "Marketing" || department == "Sales" ||
		department == "Operations" || department == "Support" {
		return true
	}
	return false
}

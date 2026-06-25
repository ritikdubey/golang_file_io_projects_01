package main

import (
	"bufio"
	"cmp"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"time"
)

type employee struct {
	EmployeeID  int
	Name        string
	Department  string
	Age         int
	Salary      int
	Experience  int
	JoiningDate time.Time
	City        string
}

func main() {

	input_file, err := os.Open("data.csv")
	if err != nil {
		log.Fatalf("error while opening file: %v", err)
	}
	defer input_file.Close()

	output_file, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error while creating file: %v", err)
	}
	defer output_file.Close()

	reader := csv.NewReader(input_file)
	writer := bufio.NewWriter(output_file)

	records, err := reader.ReadAll()

	pairs := []employee{}
	layout := "2006-01-02"

	for _, row := range records[1:] {
		employeeId, _ := strconv.Atoi(row[0])
		age, _ := strconv.Atoi(row[3])
		salary, _ := strconv.Atoi(row[4])
		experience, _ := strconv.Atoi(row[5])
		joiningDate, _ := time.Parse(layout, row[6])
		pairs = append(pairs, employee{EmployeeID: employeeId, Name: row[1], Department: row[2], Age: age,
			Salary: salary, Experience: experience, JoiningDate: joiningDate, City: row[7]})
	}

	slices.SortFunc(pairs, func(a, b employee) int {
		return cmp.Compare(a.Name, b.Name)
	})

	fmt.Fprintf(writer, "=================================================\n")
	fmt.Fprintf(writer, "DATA SORTED BY NAME \n")
	for _, r := range pairs {
		fmt.Fprintf(writer, "%v, %v, %v, %v, %v, %v, %v, %v\n",
			r.EmployeeID, r.Name, r.Department, r.Age, r.Salary, r.Experience, r.Age, r.City)
	}
	fmt.Fprintf(writer, "=================================================\n")

	slices.SortFunc(pairs, func(a, b employee) int {
		return cmp.Compare(a.Age, b.Age)
	})

	fmt.Fprintf(writer, "=================================================\n")
	fmt.Fprintf(writer, "DATA SORTED BY AGE \n")
	for _, r := range pairs {
		fmt.Fprintf(writer, "%v, %v, %v, %v, %v, %v, %v, %v\n",
			r.EmployeeID, r.Name, r.Department, r.Age, r.Salary, r.Experience, r.Age, r.City)
	}
	fmt.Fprintf(writer, "=================================================\n")

	slices.SortFunc(pairs, func(a, b employee) int {
		return cmp.Compare(a.Salary, b.Salary)
	})

	fmt.Fprintf(writer, "=================================================\n")
	fmt.Fprintf(writer, "DATA SORTED BY SALARY \n")
	for _, r := range pairs {
		fmt.Fprintf(writer, "%v, %v, %v, %v, %v, %v, %v, %v\n",
			r.EmployeeID, r.Name, r.Department, r.Age, r.Salary, r.Experience, r.Age, r.City)
	}
	fmt.Fprintf(writer, "=================================================\n")

	slices.SortFunc(pairs, func(a, b employee) int {
		return cmp.Compare(a.Department, b.Department)
	})

	fmt.Fprintf(writer, "=================================================\n")
	fmt.Fprintf(writer, "DATA SORTED BY DEPARTMENT \n")
	for _, r := range pairs {
		fmt.Fprintf(writer, "%v, %v, %v, %v, %v, %v, %v, %v\n",
			r.EmployeeID, r.Name, r.Department, r.Age, r.Salary, r.Experience, r.Age, r.City)
	}
	fmt.Fprintf(writer, "=================================================\n")

	slices.SortFunc(pairs, func(a, b employee) int {
		return cmp.Compare(a.Experience, b.Experience)
	})

	fmt.Fprintf(writer, "=================================================\n")
	fmt.Fprintf(writer, "DATA SORTED BY EXPERIENCE \n")
	for _, r := range pairs {
		fmt.Fprintf(writer, "%v, %v, %v, %v, %v, %v, %v, %v\n",
			r.EmployeeID, r.Name, r.Department, r.Age, r.Salary, r.Experience, r.Age, r.City)
	}
	fmt.Fprintf(writer, "=================================================\n")

	slices.SortFunc(pairs, func(a, b employee) int {
		return a.JoiningDate.Compare(b.JoiningDate)
	})

	fmt.Fprintf(writer, "=================================================\n")
	fmt.Fprintf(writer, "DATA SORTED BY JOINING DATE \n")
	for _, r := range pairs {
		fmt.Fprintf(writer, "%v, %v, %v, %v, %v, %v, %v, %v, %v\n",
			r.EmployeeID, r.Name, r.Department, r.Age, r.Salary, r.Experience, r.Age, r.City, r.JoiningDate)
	}
	fmt.Fprintf(writer, "=================================================\n")

	if err := writer.Flush(); err != nil {
		log.Fatalf("error while writing file: %v", err)
	}

}

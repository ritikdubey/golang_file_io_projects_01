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
)

type pair struct {
	Key   string
	Value int
}

func main() {

	input_file, err := os.Open("data.csv")
	if err != nil {
		log.Fatalf("error while reading file: %v", err)
	}
	defer input_file.Close()

	output_file, err := os.Create("report.txt")
	if err != nil {
		log.Fatalf("error while creating file: %v", err)
	}
	defer output_file.Close()

	reader := csv.NewReader(input_file)
	writer := bufio.NewWriter(output_file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("error while reading records: %v", err)
	}

	// Date,Type,Category,Description,Amount

	income := 0
	expense := 0
	incomeMap := make(map[string]int)
	expenseMap := make(map[string]int)
	expensePair := []pair{}

	for _, row := range records[1:] {
		amount, err := strconv.Atoi(row[4])
		if err != nil {
			continue
		}

		switch row[1] {
		case "Income":
			income += amount
			incomeMap[row[2]] += amount
		case "Expense":
			expense += amount
			expenseMap[row[2]] += amount
			expensePair = append(expensePair, pair{Key: row[2], Value: amount})
		}
	}

	savingRate := (float64(income-expense) / float64(income) * 100)
	status := ""

	if savingRate >= 40 {
		status = "HEALTHY"
	} else if savingRate >= 10 && savingRate < 40 {
		status = "AVERAGE"
	} else if savingRate < 10 {
		status = "CRITICAL"
	}

	slices.SortFunc(expensePair, func(a, b pair) int {
		return cmp.Compare(b.Value, a.Value)
	})

	fmt.Fprintf(writer, "=================================================\n")
	fmt.Fprintf(writer, "MONTHLY BUDGET REPORT\n")
	fmt.Fprintf(writer, "=================================================\n")
	fmt.Fprintf(writer, "Income: ₹%v\n", income)
	fmt.Fprintf(writer, "Expenses: ₹%v\n", expense)
	fmt.Fprintf(writer, "Net Savings: ₹%v\n", income-expense)
	fmt.Fprintf(writer, "Saving Rate: %.2f\n", savingRate)
	fmt.Fprintf(writer, "Status: %v", status)
	fmt.Fprintf(writer, "\n")
	fmt.Fprintf(writer, "\n")
	fmt.Fprintf(writer, "\n")

	fmt.Fprintf(writer, "=================================================\n")
	fmt.Fprintf(writer, "INCOME BREAKDOWN\n")
	fmt.Fprintf(writer, "=================================================\n")

	for key, val := range incomeMap {
		fmt.Fprintf(writer, "%v: %v\n", key, val)
	}
	fmt.Fprintf(writer, "\n")
	fmt.Fprintf(writer, "\n")
	fmt.Fprintf(writer, "\n")

	fmt.Fprintf(writer, "=================================================\n")
	fmt.Fprintf(writer, "EXPENSE BREAKDOWN\n")
	fmt.Fprintf(writer, "=================================================\n")

	for key, val := range expenseMap {
		fmt.Fprintf(writer, "%v: %v\n", key, val)
	}

	fmt.Fprintf(writer, "\n")
	fmt.Fprintf(writer, "\n")
	fmt.Fprintf(writer, "\n")

	fmt.Fprintf(writer, "=================================================\n")
	fmt.Fprintf(writer, "TOP EXPENSES\n")
	fmt.Fprintf(writer, "=================================================\n")
	fmt.Fprintf(writer, "%v: %v\n", expensePair[0].Key, expensePair[0].Value)
	fmt.Fprintf(writer, "%v: %v\n", expensePair[1].Key, expensePair[1].Value)
	fmt.Fprintf(writer, "%v: %v\n", expensePair[2].Key, expensePair[2].Value)

	if err := writer.Flush(); err != nil {
		log.Fatalf("error while writing: %v", err)
	}

}

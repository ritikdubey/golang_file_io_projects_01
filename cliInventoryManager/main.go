package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	inputFile, err := os.OpenFile("input.csv", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("error while opening file: %v", err)
	}

	reader := csv.NewReader(inputFile)
	writer := csv.NewWriter(inputFile)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("error while reading records: %v", err)
	}

	scanner := bufio.NewScanner(os.Stdin)

	productsNameMap := make(map[string]int)
	for _, row := range records[1:] {
		val, _ := strconv.Atoi(row[5])
		productsNameMap[row[1]] = productsNameMap[row[1]] + val
	}

	productsIdMap := make(map[string]int)
	for _, row := range records[1:] {
		val, _ := strconv.Atoi(row[5])
		productsIdMap[row[0]] = productsIdMap[row[0]] + val
	}

	for {
		fmt.Print("> ")
		scanner.Scan()

		input := scanner.Text()
		command := strings.TrimSpace(input)

		switch {
		case strings.Contains(command, "help"):
			fmt.Println("   inventory: lists all product names and their quantities")
			fmt.Println("   list {product id}: lists the quantity for a particular product id")
			fmt.Println("   add {product details}: adds a new product to inventory")
		case strings.Contains(command, "inventory"):

			if strings.Contains(command, "inventory value") {
				totalValue := 0.0
				for _, row := range records {
					unitPrice, _ := strconv.ParseFloat(row[4], 64)
					stockQuantity, _ := strconv.ParseFloat(row[5], 64)
					totalValue = totalValue + (unitPrice * stockQuantity)
				}
				fmt.Printf("Total Value of inventory: %.2f\n", totalValue)
				continue
			}

			if strings.Contains(command, "inventory warehouse") {
				warehouseMap := make(map[string][]string)
				for _, row := range records[1:] {
					warehouseMap[row[7]] = append(warehouseMap[row[7]], row[0])
				}
				fmt.Printf("Warehouse inventory:\n")
				for key, value := range warehouseMap {
					fmt.Printf("%v: %v\n", key, value)
				}
				continue
			}

			for key, val := range productsNameMap {
				fmt.Printf("%v : %v\n", key, val)
			}
		case strings.HasPrefix(command, "list"):
			commandSlice := strings.Split(command, " ")
			if len(commandSlice) > 2 {
				fmt.Println("Too many commands")
				continue
			}
			if _, ok := productsIdMap[commandSlice[1]]; ok {
				fmt.Printf("%v : %v\n", commandSlice[1], productsIdMap[commandSlice[1]])
			} else {
				fmt.Printf("%v not present in inventory\n", commandSlice[1])
			}
		case strings.HasPrefix(command, "add"):
			inputData := strings.Split(command[4:], ",")
			if len(inputData) < 10 {
				fmt.Println("invalid input data provided for add")
			}
			newSlice := [][]string{inputData}
			writer.WriteAll(newSlice)
			writer.Flush()
		case strings.HasPrefix(command, "update"):
			inputData := strings.Split(command, " ")
			fmt.Println(inputData)
			if len(inputData) < 3 || len(inputData) > 3 {
				fmt.Println("invalid input data for update")
			}
			for i, row := range records {
				fmt.Println(row[0], inputData[1])
				if strings.EqualFold(row[0], inputData[1]) {
					fmt.Println(records[i][5], inputData[2])
					records[i][5] = inputData[2]
					break
				}
			}
			if err = inputFile.Truncate(0); err != nil {
				log.Fatalf("error while clearing file: %v", err)
			}
			if _, err = inputFile.Seek(0, 0); err != nil {
				log.Fatalf("error while seek: %v", err)
			}
			writer.WriteAll(records)
			writer.Flush()
		case strings.HasPrefix(command, "delete"):
			inputData := strings.Split(command, " ")
			newRecords := [][]string{}
			deletedId := []string{}
			deletedId = append(deletedId, inputData[1])
			if err = inputFile.Truncate(0); err != nil {
				log.Fatalf("error while clearing file: %v", err)
			}
			if _, err = inputFile.Seek(0, 0); err != nil {
				log.Fatalf("error while seek: %v", err)
			}

			for _, row := range records {
				if slices.Contains(deletedId, row[0]) {
					continue
				} else {
					newRecords = append(newRecords, row)
				}
			}
			writer.WriteAll(newRecords)
			writer.Flush()
		case strings.HasPrefix(command, "report"):
			stockMap := make(map[string][]string)
			supplierMap := make(map[string][]string)
			reportFile, err := os.Create("report.txt")
			if err != nil {
				log.Fatalf("error while creating file: %v", err)
			}
			writer2 := bufio.NewWriter(reportFile)
			for _, row := range records[1:] {
				stockMap[row[9]] = append(stockMap[row[9]], row[1])
				supplierMap[row[3]] = append(supplierMap[row[3]], row[0]+"_"+row[1])
			}
			for key, value := range stockMap {
				fmt.Fprintf(writer2, "%v:\n", key)
				for _, item := range value {
					fmt.Fprintf(writer2, "%v\n", item)
				}
				fmt.Fprintf(writer2, "\n")
			}
			fmt.Fprintf(writer2, "*******************************************************************\n\n")
			for key, value := range supplierMap {
				fmt.Fprintf(writer2, "%v:\n", key)
				for _, item := range value {
					fmt.Fprintf(writer2, "%v\n", item)
				}
				fmt.Fprintf(writer2, "\n")
			}
			if err := writer2.Flush(); err != nil {
				log.Fatalf("error while writing: %v\n", err)
			}
		case strings.HasPrefix(command, "get"):
			inputData := strings.Split(command, " ")
			for _, row := range records {
				if row[0] == inputData[1] {
					fmt.Printf("%v Info\n", row[1])
					fmt.Printf("Price: %v Stock Quantity: %v\n", row[4], row[5])
					fmt.Printf("Warehouse: %v Last Restocked: %v\n", row[7], row[8])
					break
				}
			}
		case strings.EqualFold(command, "exit"):
			return
		default:
			fmt.Println("Unknown command")
		}

		if err := scanner.Err(); err != nil {
			log.Printf("Error encountered during scanning: %v", err)
		}
	}

}

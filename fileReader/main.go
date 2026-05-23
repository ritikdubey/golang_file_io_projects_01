package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Method 1
	// fileData, err := os.ReadFile("test.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(fileData))

	// Method 2
	// file, err := os.Open("test.txt")
	// if err != nil {
	// 	fmt.Println("Error while opening: ", err)
	// 	return
	// }
	// defer file.Close()

	// scanner := bufio.NewScanner(file)

	// count := 1
	// for scanner.Scan() {
	// 	fmt.Println(count)
	// 	count += 1
	// 	line := scanner.Text()
	// 	fmt.Println(line)
	// }

	// if err := scanner.Err(); err != nil {
	// 	fmt.Println("Error during scan: ", err)
	// }

	capitals := make(map[string]string)
	capitals["India"] = "New Delhi"
	capitals["Australia"] = "Canberra"
	capitals["USA"] = "Washington DC"
	capitals["Sri Lanka"] = "Colombo"
	capitals["UK"] = "London"
	capitals["Japan"] = "Tokyo"
	capitals["Vietnam"] = "Ho Chin Minh City"
	capitals["Russia"] = "Moscow"
	capitals["China"] = "Bejing"
	capitals["Sweden"] = "Stockholm"

	input_file, err := os.Open("test.txt")

	if err != nil {
		fmt.Println("error while reading file: ", err)
		return
	}
	defer input_file.Close()

	output_file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("error while creating file: ", err)
		return
	}
	defer output_file.Close()

	scanner := bufio.NewScanner(input_file)
	writer := bufio.NewWriter(output_file)

	for scanner.Scan() {
		line := scanner.Text()
		//_, err := writer.WriteString(fmt.Fprintf(writer, "%s - %s\n", country, capital))
		fmt.Fprintf(writer, "%s - %s\n", line, capitals[line])
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error during scan: ", err)
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer:", err)
	}

}

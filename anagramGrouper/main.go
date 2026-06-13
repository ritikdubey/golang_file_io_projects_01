package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {

	inputBytes, err := os.ReadFile("dictionary.json")

	if err != nil {
		log.Fatal("error while reading file: ", err)
	}

	output_file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal("error while creating file: ", err)
	}

	defer output_file.Close()

	writer := bufio.NewWriter(output_file)

	var data map[string]string

	if err := json.Unmarshal(inputBytes, &data); err != nil {
		log.Fatal("Failed to parse json: ", err)
	}

	// fmt.Println(data["butcher"])

	resultMap := make(map[string][]string)

	for key := range data {

		runes := []rune(key)
		slices.Sort(runes)
		sortedKey := string(runes)

		// if _, ok := resultMap[sortedKey]; ok {
		// 	resultMap[sortedKey] = append(resultMap[sortedKey], key)
		// } else {
		// 	resultMap[sortedKey] = []string{key}
		// }

		resultMap[sortedKey] = append(resultMap[sortedKey], key)

	}

	// fmt.Println(resultMap)

	for key, value := range resultMap {
		if len(value) > 1 {
			fmt.Fprintf(writer, "%s : %s\n", key, value)
		}
	}

	err = writer.Flush()
	if err != nil {
		log.Fatal("error while flush")
	}

}

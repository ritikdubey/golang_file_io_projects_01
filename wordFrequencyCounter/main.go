package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	resultMap := make(map[string]int)

	input_file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("error while reading file: ", err)
	}

	defer input_file.Close()

	output_file, err := os.Create("output_file.txt")

	if err != nil {
		fmt.Println("error while creating file: ", err)
	}

	defer output_file.Close()

	scanner := bufio.NewScanner(input_file)

	writer := bufio.NewWriter(output_file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ToLower(line)

		checkSpecialCharacters(&resultMap, &line)

		wordsList := strings.Split(line, " ")

		for _, v := range wordsList {
			if _, ok := resultMap[v]; ok {
				resultMap[v]++
			} else {
				resultMap[v] = 1
			}
		}

	}

	keys := make([]string, 0, len(resultMap))
	for k := range resultMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println(resultMap)

	for key, value := range resultMap {
		fmt.Fprintf(writer, "%s : %v\n", key, value)
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer:", err)
	}

}

func checkSpecialCharacters(resultMap *map[string]int, line *string) {

	commaCount := strings.Count(*line, ",")
	fullStopCount := strings.Count(*line, ".")
	ampersandCount := strings.Count(*line, "&")
	singleQuotesCount := strings.Count(*line, "'")
	doubleQuotesCount := strings.Count(*line, "\"")
	questionMarkCount := strings.Count(*line, "?")
	hyphenCount := strings.Count(*line, "-")

	specialCharacterCount := commaCount + fullStopCount + ampersandCount + singleQuotesCount + doubleQuotesCount +
		questionMarkCount + hyphenCount

	// fmt.Println("specialCharacterCount: ", specialCharacterCount)

	if specialCharacterCount > 0 {
		*line = strings.ReplaceAll(*line, ",", "")
		*line = strings.ReplaceAll(*line, ".", "")
		*line = strings.ReplaceAll(*line, "&", "")
		*line = strings.ReplaceAll(*line, "'", "")
		*line = strings.ReplaceAll(*line, "\"", "")
		*line = strings.ReplaceAll(*line, "?", "")
		*line = strings.ReplaceAll(*line, "-", "")

		// fmt.Println(*line)

		if _, ok := (*resultMap)["#specialCharacterCount"]; ok {
			(*resultMap)["#specialCharacterCount"] = (*resultMap)["#specialCharacterCount"] + specialCharacterCount
		} else {
			(*resultMap)["#specialCharacterCount"] = specialCharacterCount
		}
	}

}

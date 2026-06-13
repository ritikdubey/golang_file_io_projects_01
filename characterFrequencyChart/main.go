package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {

	output, err1 := os.Create("output.txt")
	if err1 != nil {
		log.Fatalf("error while creating a new file: %s", err1)
	}

	writer := bufio.NewWriter(output)

	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("error while reading file: %s", err)
	}

	content := string(data)

	contentSlice := []rune(content)

	charMap := make(map[string]int)

	for _, v := range contentSlice {
		if string(v) == " " {
			continue
		}
		if _, ok := charMap[strings.ToLower(string(v))]; ok {
			charMap[strings.ToLower(string(v))] = charMap[strings.ToLower(string(v))] + 1
		} else {
			charMap[strings.ToLower(string(v))] = 1
		}
	}

	keys := []string{}

	for key := range charMap {
		keys = append(keys, key)
	}

	slices.SortFunc(keys, func(a, b string) int {
		return charMap[b] - charMap[a]
	})

	for _, v := range keys {
		fmt.Fprintf(writer, "%s - ", v)
		fmt.Fprintf(writer, "%s", strings.Repeat("*", charMap[v]))
		fmt.Fprintf(writer, "\n")
	}

	if err = writer.Flush(); err != nil {
		log.Fatal("error while flush")
	}

}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("error while reading file")
	}

	output, err := os.Create("output.txt")
	if err != nil {
		log.Fatal("error while creating file")
	}

	defer output.Close()

	writer := bufio.NewWriter(output)

	text := string(data)

	paragraphs := strings.Split(text, "\n")
	fmt.Println(len(paragraphs))

	for _, p := range paragraphs {
		cleanPara := strings.TrimSpace(p)
		lines := strings.Split(cleanPara, ".")
		if len(cleanPara) > 1 {
			fmt.Fprintf(writer, "%s.\n", lines[0])
		}
	}

	if err = writer.Flush(); err != nil {
		log.Fatal("error while writing")
	}

}

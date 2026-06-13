package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	matches, err := filepath.Glob("./*.txt")
	if err != nil {
		log.Fatal("error while reading")
	}

	for _, path := range matches {
		content, _ := os.ReadFile(path)
		updatedContent := strings.ReplaceAll(string(content), "NOTHING", "apples")

		err = os.WriteFile(path, []byte(updatedContent), 0644)
		if err != nil {
			log.Fatalf("Failed to write file: %s", err)
		}
	}

}

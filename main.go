package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: hfind <search-term>")
		os.Exit(1)
	}

	searchTerm := os.Args[1]

	fmt.Println("Searching for:", searchTerm)
	fmt.Println("___")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error finding home directory:", err)
	}

	historyFilePath := filepath.Join(homeDir, ".zsh_history")

	file, err := os.Open(historyFilePath)
	if err != nil {
		log.Fatalf("Error opening history file at %s: %v", historyFilePath, err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, searchTerm) {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading history file:", err)
	}
}

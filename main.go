package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
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

	if err := runSearch(searchTerm); err != nil {
		log.Fatal(err)
	}
}

func runSearch(searchTerm string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error finding home directory: %w", err)
	}

	historyFilePath := filepath.Join(homeDir, ".zsh_history")
	file, err := os.Open(historyFilePath)
	if err != nil {
		return fmt.Errorf("error opening history file at %s: %w", historyFilePath, err)
	}
	defer file.Close()

	red := color.New(color.FgRed, color.Bold).SprintFunc()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, searchTerm) {
			coloredLine := strings.ReplaceAll(line, searchTerm, red(searchTerm))
			fmt.Println(coloredLine)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading history file: %w", err)
	}

	return nil
}

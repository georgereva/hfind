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
	historyFilePath, err := getHistoryFilePath()
	if err != nil {
		return err
	}

	file, err := os.Open(historyFilePath)
	if err != nil {
		return fmt.Errorf("could not open history file at %s: %w", historyFilePath, err)
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

func getHistoryFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not find home directory: %w", err)
	}

	shellPath := os.Getenv("SHELL")
	if shellPath == "" {
		return filepath.Join(homeDir, ".bash_history"), nil
	}

	shell := filepath.Base(shellPath)

	switch shell {
	case "zsh":
		if histFile := os.Getenv("HISTFILE"); histFile != "" {
			return histFile, nil
		}

		return filepath.Join(homeDir, ".zsh_history"), nil
	case "bash":
		return filepath.Join(homeDir, ".bash_history"), nil
	case "sh":
		return filepath.Join(homeDir, ".sh_history"), nil
	default:
		return "", fmt.Errorf("unsupported shell: %s. Only bash, sh, zsh are supported", shell)
	}
}

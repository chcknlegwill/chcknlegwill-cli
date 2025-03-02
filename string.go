package main

//work on highlighting it in the terminal
//propably have to use some other binary.

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

func isReadableFile(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()

	buf := make([]byte, 512)
	n, err := file.Read(buf)
	if err != nil {
		return false
	}
	//checks if valid utf-8 & not a binary file (contains null value)
	return utf8.Valid(buf[:n]) && !strings.Contains(string(buf[:n]), "\x00")
}

func searchFiles(searchStr string) error {
	found := false
	// Walk through the current directory to find files
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Only skip hidden files and directories (names starting with .) if they are not the root directory "."
		if path != "." && strings.HasPrefix(info.Name(), ".") {
			if info.IsDir() {
				//fmt.Printf("Skipping hidden dir: %s\n", path)
				return filepath.SkipDir // Skip entire hidden directories like .git
			}
			//fmt.Printf("Skipping hidden file: '%s'\n", path)
			return nil // Skip hidden files
		}

		// Only process regular files (not directories)
		if !info.IsDir() {
			if !isReadableFile(path) {
				//				fmt.Printf("Skipping unreadable file: %s\n", path)
				return nil
			}
			//fmt.Printf("Processing file: %s\n", path)
			line, lineNumber, err := Readln(searchStr, path) // searchStr is the search string, path is the file path
			if err != nil && !strings.Contains(err.Error(), "not found") {
				return fmt.Errorf("error reading %s: %v", path, err)
			}
			if line != "" {
				found = true
				trimmed := strings.TrimSpace(line)
				fmt.Printf("Found '%s' in %s on line %d: %s\n", searchStr, path, lineNumber, trimmed)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	// If no matches were found, inform the user
	if !found {
		fmt.Printf("String '%s' not found in any files.\n", searchStr)
	}
	return nil
}

func Readln(searchStr, path string) (string, int, error) {
	content, err := os.ReadFile(path) // Read the file at the given path
	if err != nil {
		return "", 0, err
	}

	if strings.Contains(strings.ToLower(string(content)), strings.ToLower(searchStr)) { // Search for searchStr in content
		lines := strings.Split(string(content), "\n")
		for i, line := range lines {
			if strings.Contains(strings.ToLower(line), strings.ToLower(searchStr)) {
				return line, i + 1, nil // Return the line, line number (1-based), and no error
			}
		}
	}
	return "", 0, fmt.Errorf("string '%s' not found in file %s", searchStr, path)
}

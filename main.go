package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//chcknlegwill-cli v1.0.0

func main() {
	// Collect and clean command-line arguments
	input := make([]string, 0, len(os.Args)-1)
	for _, arg := range os.Args[1:] {
		trimmed := strings.TrimSpace(arg)
		if trimmed != "" {
			input = append(input, trimmed)
		}
	}

	if len(input) == 0 {
		fmt.Println("chcknlegwill-cli v1.0.0")
	}

	searchString := flag.String("f", "", "Search for a string")
	flag.Parse()

	if *searchString == "" {
		args := make([]string, 0, len(os.Args))
		for _, arg := range os.Args[1:] {
			trimmed := strings.TrimSpace(arg)
			if trimmed != "" {
				args = append(args, trimmed)
			}
		}
	}

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		if strings.Contains(strings.ToLower(string(content)), strings.ToLower(*searchString)) {
			fmt.Printf("Found '%s' in: %s\n", *searchString, path)
		}
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error walking the path: %v\n", err)
		os.Exit(1)
	}

}

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func searchFiles(searchStr string) error {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}

		if strings.Contains(strings.ToLower(string(content)), strings.ToLower(searchStr)) {
			fmt.Printf("Found '%s' in: %s \n", searchStr, path)
			//fmt.Printf("On line: x | [line content]")
			//make it show the line number and the rest strings it found
			//e.g. if it's on line 8: (with) "bruh" -> it will show you the whole thing
		}
		return nil
	})
	return err
}

package main

//work on highlighting it in the terminal
//propably have to use some other binary.

// Get the full path (include current Dir in the output)

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/ttacon/chalk"
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

func searchFiles(searchStr string, dir string) error {
	//red := chalk.Red.NewStyle().WithBackground(chalk.Red)
	green := chalk.Green.NewStyle()
	found := false

	//Walk through the current directory to find files
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//skip hidden files and directories (names starting with .) if they are not the root directory "."
		if path != "." && strings.HasPrefix(info.Name(), ".") {
			if info.IsDir() {
				//fmt.Printf(red.Style("Skipping hidden dir: %s\n"), path)
				return filepath.SkipDir // Skip entire hidden directories like .git
			}
			//red style may be too aggressive (reserve for errors...)
			//fmt.Printf(red.Style("Skipping hidden file: '%s'\n"), path)
			return nil // Skip hidden files
		}

		//only process regular files (not directories)
		if !info.IsDir() {
			if !isReadableFile(path) {
				//fmt.Printf("Skipping unreadable file: %s\n", path)
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
				//colour in the found string green so its easier to read
				fmt.Printf(green.Style("Found")+" '%s' in file: %s on line %d: "+green.Style("%s\n"), searchStr, path, lineNumber, trimmed)
				//don't use commas unless you need to ^ concatenation works well.

				//make an improved -v (verbose flag) for all options to give the user more info about what the cli is doing.
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	//if no matches found
	if !found {
		fmt.Printf("String '%s' not found in any files in %s.\n", searchStr, dir)
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
				return line, i + 1, nil //return the line, line number (1-based), and no error
			}
		}
	}
	return "", 0, fmt.Errorf("string '%s' not found in file %s", searchStr, path)
}

//for the --verbose flag:

func listDirectoryStructure(rootPath string) error {
	fmt.Println("Listing directory structure from root (/)...")

	yellow := chalk.Yellow.NewStyle()

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			//skip directories user does not have permission to access
			if os.IsPermission(err) {
				fmt.Fprintf(os.Stderr, yellow.Style("Warning:")+" Permission denied for %s\n", path)
				return filepath.SkipDir
			}
			return err
		}

		//want verbose to uncover hidden files (prefixed with a "." e.g. .gitignore)

		// Print the path with indentation to show hierarchy
		prefix := strings.Repeat("  ", max(0, strings.Count(path, string(os.PathSeparator))-1))
		//improve the output structure instead of icons, show the structure with the file extensions (maybe)
		/*
			e.g.
			for Directory: Documents/notes.txt, Documents/notes/notes2.txt
			equals: Documents/
									|- note.txt
									|- notes/
											|- note2.txt
		*/
		if info.IsDir() {
			fmt.Printf("%sfolder: %s/\n", prefix, info.Name())
		} else {
			fmt.Printf("%sfile:  %s\n", prefix, info.Name())
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to walk directory: %v", err)
	}
	return nil
}

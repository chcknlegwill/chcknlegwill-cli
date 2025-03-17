package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

//try to keep imports to a minumum as it just creates more headaches
//^ like node (node_modules bigger than the universe)

// chcknlegwill-cli v1.0.4

func init() {

}

func main() {
	//define CLI flags ("-h", "--help", "-f")
	//working functions
	Dir := ""
	searchForString := pflag.StringP("search", "f", "", "Search for a string in files and folders recursively.")
	list := pflag.BoolP("list", "l", false, "List the entire directory from the root (/).")

	//below are helper functions
	version := pflag.BoolP("version", "v", false, "Show the version of the program")
	help := pflag.BoolP("help", "h", false, "Show this help message.")
	pflag.CommandLine.SortFlags = false
	pflag.Parse()

	//working - add a verbose flag that outputs skipped files
	//as well as optional thorough search that includes all files prefixed with "."
	//e.g. .DS_Store .gitignore etc...
	//red := chalk.Red.NewStyle().WithForeground(chalk.Red)
	//fmt.Print(red.Style("Bruh moment"))

	//show help if requested
	if *help || (pflag.NFlag() == 0) {
		fmt.Println("Usage:\nchcknlegwill-cli -f <string> | -l | -h | | --help")
		fmt.Println("-f, --search <string>      Search for a string in files and folders recursivley.")
		fmt.Println("-l, --list   <directory>   List the entire directory from current directory.")
		fmt.Println("-h, --help	    	   Show this help message.")
		//change to version thing
		fmt.Println("-v, --version	    	   Show the version of the program.")
		return
	}

	//show the version of the program
	if *version {
		fmt.Println("v1.0.4")
		return
	}

	//check if -f is provided but no value is given
	iFlag := pflag.Lookup("f")
	if iFlag != nil && iFlag.Changed && *searchForString == "" {
		fmt.Fprintf(os.Stderr, "Error: The -f flag requires a search string (e.g., -f keyword)\n")
		os.Exit(1)
	}

	// If -f is provided with a string, trigger search functionality

	//TODO: Update with searching in specifiec directories e.g. if you are in one dir
	//and want to search a different folder you would write:
	//chcknlegwill-cli -f <string_to_search> <directory_to_search>
	//may need to change the signature of the flags or create a custom one or smth like that to get
	//multiple arguments from the cli as it's not working with any extra strings, not even crashing
	if *searchForString != "" {
		if Dir != "" {
			return
		}
		err := searchFiles(*searchForString) // Pass the search string to searchFiles

		if err != nil {
			//fmt.Fprintf(os.Stderr, red.Style("Error during search: %v\n"), err)
			os.Exit(1)
		}
	}

	if *list {
		err := listDirectoryStructure(".")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error listing directory structure: %v\n", err)
			os.Exit(1)
		}
		return
	}
}

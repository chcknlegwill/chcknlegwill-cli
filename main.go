package main

import (
    "fmt"
    "os"
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

    /*
    // Handle different cases based on the number of input 
    if len(input) == 0 {
        fmt.Println("chcknlegwill-cli v1.0.0")
    } else if len(input) == 1 {
        fmt.Printf("Hello, %s!\n", input[0])
    } else if len(input) == 2 {
        fmt.Printf("Hello, %s and %s!\n", input[0], input[1])
    } else {
        last := input[len(input)-1]
        others := strings.Join(input[:len(input)-1], ", ")
        fmt.Printf("Hello, %s, and %s!\n", others, last)
    }
    */
}

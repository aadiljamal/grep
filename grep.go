package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "regexp"
)

func main() {
    // Define command-line flags
    pattern := flag.String("pattern", "", "the regular expression pattern to search for")
    filename := flag.String("file", "", "the name of the file to search in")

    flag.Parse()

    // Check that both flags are provided
    if *pattern == "" || *filename == "" {
        fmt.Println("Usage: grep -pattern=<pattern> -file=<filename>")
        os.Exit(1)
    }

    // Compile the regular expression pattern
    re := regexp.MustCompile(*pattern)

    // Open the file for reading
    file, err := os.Open(*filename)
    if err != nil {
        fmt.Println("Error opening file:", err)
        os.Exit(1)
    }
    defer file.Close()

    // Create a scanner to read the file line by line
    scanner := bufio.NewScanner(file)

    // Iterate over each line of the file
    for scanner.Scan() {
        line := scanner.Text()

        // Check if the line matches the pattern
        if re.MatchString(line) {
            fmt.Println(line)
        }
    }

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        fmt.Println("Error scanning file:", err)
        os.Exit(1)
    }
}


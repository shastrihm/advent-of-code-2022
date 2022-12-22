package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func Part_one(input_file string) int {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()

    }
    return 0
}

func Part_two(input_file string) int {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()
    }
    return 0
}

func main() {
    input := os.Args[1]
    fmt.Println(Part_one(input))
    fmt.Println(Part_two(input))
}

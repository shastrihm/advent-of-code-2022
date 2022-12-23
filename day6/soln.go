package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func Part_one(input_file string) int {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    // Change this to 14 for part 2
    windowSize := 4
    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()
        for pos, _ := range line {
            if len(line) > pos + windowSize {
                substr := line[pos:pos + windowSize]
                pass := true
                for _, ch := range substr {
                    pass = pass && strings.Count(substr, string(ch)) == 1
                }
                if pass {
                    return windowSize + pos
                }
            }
        }
    }
    return -1
}

// Part two is identical to part 1 except for windowSize

// func Part_two(input_file string) int {
//     file, err := os.Open(input_file)
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer file.Close()
//
//     scanner := bufio.NewScanner(file)
//     // Read input line by line
//     for scanner.Scan() {
//         line := scanner.Text()
//         // code here
//     }
//     return 0
// }

func main() {
    input := os.Args[1]
    fmt.Println(Part_one(input))
    //fmt.Println(Part_two(input))
}

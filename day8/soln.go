package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)


func Part_one(input_file string) int {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    treeGrid := make([][]int, 0)
    // Read input line by line into 2d slice
    for scanner.Scan() {
        line := scanner.Text()
        row := make([]int, len(line))
        for pos, digit := range line {
            num, err := strconv.Atoi(string(digit))
            if err != nil {
                log.Fatal(err)
            }
            row[pos] = num
        }
        treeGrid = append(treeGrid, row)
    }

    // Actually solve the problem
    fmt.Println(treeGrid)
    return 0
}

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

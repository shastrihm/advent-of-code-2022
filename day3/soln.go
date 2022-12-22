package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func Priority(item rune) int {
    if item >= 97 {
        return int(item) - 96
    } else {
        return int(item) - 38
    }
}

func Part_one(input_file string) int {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    sumPriorities := 0
    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()
        contains := map[rune]int{}

        length := len(line)
        for i, ch := range line {
            if i < length/2 {
                contains[ch] = 1
            } else {
                _, ok := contains[ch]
                if ok {
                    sumPriorities += Priority(ch)
                    break
                }
            }
        }
    }
    return sumPriorities
}

func Part_two(input_file string) int {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // Read input line by line
    counter := 0
    contains := map[rune]int{}
    sumPriorities := 0
    for scanner.Scan() {
        line := scanner.Text()
        step := counter % 3
        for _, ch := range line {
            v, _ := contains[ch]
            if v == step {
                contains[ch] += 1
                if contains[ch] == 3 {
                    contains = map[rune]int{}
                    sumPriorities += Priority(ch)
                    break
                }
            }
        }

        counter = counter + 1
    }
    return sumPriorities
}

func main() {
    input := os.Args[1]
    fmt.Println(Part_one(input))
    fmt.Println(Part_two(input))
}

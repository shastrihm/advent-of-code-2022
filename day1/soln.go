package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "sort"
)

func Part_one(input_file string) int {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // Read input line by line
    best := 0
    currentSum := 0
    for scanner.Scan() {
        line := scanner.Text()

        if line == "" {
            if currentSum > best {
                best = currentSum
            }
            currentSum = 0
            continue
        }

        cal, err := strconv.Atoi(line)
        if err != nil {
            log.Fatal(err)
        }

        currentSum = currentSum + cal
    }

    // do last elf manually due to EOF
    if currentSum > best {
        best = currentSum
    }

    return best
}

func Part_two(input_file string) int {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // Read input line by line
    best := []int{0, 0, 0, 0}
    currentSum := 0
    for scanner.Scan() {
        line := scanner.Text()

        if line == "" {
            best[0] = currentSum
            // sort best and throw away the lowest since we only wamt top 3
            sort.Slice(best, func(i, j int) bool {
                return best[i] < best[j]
            })
            best[0] = 0

            currentSum = 0
            continue
        }

        cal, err := strconv.Atoi(line)
        if err != nil {
            log.Fatal(err)
        }
        currentSum = currentSum + cal
    }

    // do the last elf manually due to EOF
    best[0] = currentSum
    sort.Slice(best, func(i, j int) bool {
        return best[i] < best[j]
    })
    best[0] = 0

    // Sum the values of the best 3 elves
    return best[1] + best[2] + best[3]
}

func main() {
    input := os.Args[1]
    fmt.Println(Part_one(input))
    fmt.Println(Part_two(input))
}

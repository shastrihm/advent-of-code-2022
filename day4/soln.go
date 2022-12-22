package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

func ParsePairing(pair string) [][]int {
    pairing := strings.Split(pair, ",")

    elf1 := strings.Split(pairing[0], "-")
    elf1_a, err := strconv.Atoi(elf1[0])
    if err != nil {
        log.Fatal(err)
    }
    elf1_b, err := strconv.Atoi(elf1[1])
    if err != nil {
        log.Fatal(err)
    }

    elf2 := strings.Split(pairing[1], "-")
    elf2_a, err := strconv.Atoi(elf2[0])
    if err != nil {
        log.Fatal(err)
    }
    elf2_b, err := strconv.Atoi(elf2[1])
    if err != nil {
        log.Fatal(err)
    }

    return [][]int{
                {elf1_a, elf1_b},
                {elf2_a, elf2_b},
    }
}

func Part_one(input_file string) int {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    fully_contained := 0
    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()

        elf := ParsePairing(line)

        if (elf[0][0] <= elf[1][0] && elf[0][1] >= elf[1][1]) ||
            (elf[1][0] <= elf[0][0] && elf[1][1] >= elf[0][1]) {
            fully_contained += 1
        }

    }
    return fully_contained
}

func Part_two(input_file string) int {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    overlaps := 0
    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()
        elf := ParsePairing(line)

        if (elf[0][0] <= elf[1][0] && elf[1][0] <= elf[0][1]) ||
            (elf[1][0] <= elf[0][0] && elf[0][0] <= elf[1][1]) {
            overlaps += 1
        }
    }
    return overlaps
}

func main() {
    input := os.Args[1]
    fmt.Println(Part_one(input))
    fmt.Println(Part_two(input))
}

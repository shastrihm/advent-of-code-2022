package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)
const win = 6
const draw = 3
const lose = 0

const rock = 1
const paper = 2
const scissors = 3

func outcomeMap(part int) map[string]int {
    if part == 1 {
        return map[string]int{
            "A X": draw + rock,
            "A Y": win + paper,
            "A Z": lose + scissors,
            "B X": lose + rock,
            "B Y": draw + paper,
            "B Z": win + scissors,
            "C X": win + rock,
            "C Y": lose + paper,
            "C Z": draw + scissors,
        }
    } else {
        return map[string]int{
            "A X": lose + scissors,
            "A Y": draw + rock,
            "A Z": win + paper,
            "B X": lose + rock,
            "B Y": draw + paper,
            "B Z": win + scissors,
            "C X": lose + paper,
            "C Y": draw + scissors,
            "C Z": win + rock,
        }
    }
}


func Part_one(input_file string) int {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    totalScore := 0
    score := outcomeMap(1)
    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()
        totalScore = totalScore + score[line]
    }
    return totalScore
}

func Part_two(input_file string) int {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    totalScore := 0
    score := outcomeMap(2)
    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()
        totalScore = totalScore + score[line]
    }
    return totalScore
}

func main() {
    input := os.Args[1]
    fmt.Println(Part_one(input))
    fmt.Println(Part_two(input))
}

package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func transpose(a [][]int) [][]int {
    newArr := make([][]int, len(a))
    for i := 0; i < len(a); i++ {
        for j := 0; j < len(a[0]); j++ {
            newArr[j] = append(newArr[j], a[i][j])
        }
    }
    return newArr
}


// returns true if val is the unique maximum in arr
func UniqueMax(arr []int, val int) bool {
    apps := 0
    for _, n := range arr {
        if n > val {
            return false
        }
        if n == val {
            apps += 1
        }
    }
    return apps == 1
}

// if dir = 1, counts viewing distance from left side of arr
// if dir = -1, counts viewing distance from right side of arr
func ViewingDistance(arr []int, dir int) int {
    viewDist := 0
    if dir == 1 {
        tree := arr[0]
        for i := 1; i < len(arr); i++ {
            if arr[i] >= tree {
                viewDist += 1
                return viewDist
            }
            viewDist += 1
        }
    } else {
        tree := arr[len(arr)-1]
        for i := len(arr) - 2; i > -1; i-- {
            if arr[i] >= tree {
                viewDist += 1
                return viewDist
            }
            viewDist += 1
        }
    }

    return viewDist
}

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
    treeGridTransposed := transpose(treeGrid)
    visible := 0
    for i := 0; i < len(treeGrid); i++ {
        for j := 0; j < len(treeGrid[0]); j++ {
            tree := treeGrid[i][j]
            row := treeGrid[i]
            col := treeGridTransposed[j]
            if UniqueMax(row[0:j+1], tree) ||
                UniqueMax(row[j:len(row)], tree) ||
                UniqueMax(col[0:i+1], tree) ||
                UniqueMax(col[i:len(col)], tree) {
                    visible += 1
            }
        }
    }
    return visible
}

func Part_two(input_file string) int {
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
    // solve problem
    treeGridTransposed := transpose(treeGrid)
    bestScenicScore := 0
    for i := 0; i < len(treeGrid); i++ {
        for j := 0; j < len(treeGrid[0]); j++ {
            row := treeGrid[i]
            col := treeGridTransposed[j]
            scenicScore := ViewingDistance(row[0:j+1], -1) *
                ViewingDistance(row[j:len(row)], 1) *
                ViewingDistance(col[0:i+1], -1) *
                ViewingDistance(col[i:len(col)], 1)

            if scenicScore > bestScenicScore {
                bestScenicScore = scenicScore
            }
        }
    }

    return bestScenicScore
}

func main() {
    input := os.Args[1]
    fmt.Println(Part_one(input))
    fmt.Println(Part_two(input))
}

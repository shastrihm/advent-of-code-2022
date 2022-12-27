package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

type Pos [2]int // row, col

func getLegalMoves(grid [][]rune, currRow int, currCol int) []Pos {
    currElev := grid[currRow][currCol]
    if currElev== 'S' {
        currElev = 'a'
    }

    possibleMoves := []Pos{
        {currRow + 1, currCol},
        {currRow - 1, currCol},
        {currRow, currCol + 1},
        {currRow, currCol - 1},
    }

    var legalMoves []Pos
    var elev rune

    for _, m := range possibleMoves {
        newRow := m[0]
        newCol := m[1]
        if newRow <= len(grid) - 1 &&
            newRow >= 0 &&
            newCol <= len(grid[0]) - 1 &&
            newCol >= 0 {
                elev = grid[newRow][newCol]
                if elev == 'E' {
                    elev = 'z'
                }
                if elev <= currElev + 1 {
                    legalMoves = append(legalMoves, Pos{newRow, newCol})
                }
            }
    }
    return legalMoves
}

func computePathLength(backtrack map[Pos]Pos, start, end Pos) int {
    curr := end
    length := 0
    for curr != start {
        curr = backtrack[curr]
        length += 1
    }
    return length
}

// Breadth-first search
// returns length of path from start ('S') to finish ('E')
// returns -1 if end not found
func BFS(grid [][]rune, start Pos) int {
    queue := []Pos{start}
    visited := map[Pos]bool{}
    backtrack := map[Pos]Pos{}

    for len(queue) > 0 {
        next := queue[0]
        queue = queue[1:]
        visited[next] = true
        if grid[next[0]][next[1]] == 'E' {
            return computePathLength(backtrack, start, next)
        }
        for _, move := range getLegalMoves(grid, next[0], next[1]) {
            _, ok := visited[move]
            if !ok {
                queue = append(queue, move)
                backtrack[move] = next
            }
            visited[move] = true
        }
    }
    return -1
}

func Part_one(input_file string) int {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    grid := [][]rune{}
    var start Pos
    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()
        row := []rune{}
        for i, e := range line {
            row = append(row, e)
            if e == 'S' {
                start = Pos{len(grid), i}
            }
        }
        grid = append(grid, row)
    }
    return BFS(grid, start)
}

func Part_two(input_file string) int {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    grid := [][]rune{}
    starts := []Pos{}
    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()
        row := []rune{}
        for i, e := range line {
            row = append(row, e)
            if e == 'S' || e == 'a' {
                starts = append(starts, Pos{len(grid), i})
            }
        }
        grid = append(grid, row)
    }
    best := -1
    for _, s := range starts {
        res := BFS(grid, s)
        if (best > res || best == -1) && res != -1 {
            best = res
        }
    }
    return best
}

func main() {
    input := os.Args[1]
    fmt.Println(Part_one(input))
    fmt.Println(Part_two(input))
}

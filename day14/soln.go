package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

type Pos [2]int

type Simulation struct {
    yLimit int
    caveMap map[Pos]rune
    source Pos
    sandUnitsAtRest int
    currSandPos Pos
}

func (s *Simulation) GetNextPos(currSand Pos) Pos {
    nextPos := Pos{currSand[0], currSand[1] + 1}
    if _, ok := s.caveMap[nextPos]; !ok && nextPos[1] < s.yLimit + 2 {
        return nextPos
    }
    tryPosLeft := Pos{nextPos[0] - 1, nextPos[1]}
    if _, ok := s.caveMap[tryPosLeft]; !ok && tryPosLeft[1] < s.yLimit + 2 {
        return tryPosLeft
    }
    tryPosRight := Pos{nextPos[0] + 1, nextPos[1]}
    if _, ok := s.caveMap[tryPosRight]; !ok && tryPosRight[1] < s.yLimit + 2 {
        return tryPosRight
    }
    // otherwise sand must be at at rest
    return currSand
}

// returns true once currentSand surpasses yLimit,
// and thus can never come to rest (part 1)
//
// returns true once currentSand comes to rest at
// source (part 2)
func (s *Simulation) Step(part int) bool {
    newPos := s.GetNextPos(s.currSandPos)
    if newPos == s.currSandPos {
        // at rest
        s.caveMap[s.currSandPos] = 'o'
        s.sandUnitsAtRest++
        if newPos == s.source && part == 2 {
            return true
        }
        // new grain
        s.currSandPos = s.source
        return false
    } else if newPos[1] == s.yLimit && part == 1 {
        return true
    }
    s.currSandPos[0] = newPos[0]
    s.currSandPos[1] = newPos[1]
    return false
}

func initSimFromInput(input_file string) *Simulation {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    caveMap := map[Pos]rune{}
    limit := 0
    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()
        // code here
        coords := strings.Split(line, " -> ")
        prev := Pos{-1,-1}
        for _, coord := range coords {
            xy := strings.Split(coord, ",")
            col, err := strconv.Atoi(xy[0])
            if err != nil {
                log.Fatal(err)
            }

            row, err := strconv.Atoi(xy[1])
            if err != nil {
                log.Fatal(err)
            }

            // set rocks in cave
            if prev[0] == col {
                if prev[1] < row {
                    for i := prev[1]; i <= row; i++ {
                        caveMap[Pos{col, i}] = '#'
                    }
                } else {
                    for i := prev[1]; i >= row; i-- {
                        caveMap[Pos{col, i}] = '#'
                    }
                }
            } else if prev[1] == row {
                if prev[0] < col {
                    for i := prev[0]; i <= col; i++ {
                        caveMap[Pos{i, row}] = '#'
                    }
                } else {
                    for i := prev[0]; i >= col; i-- {
                        caveMap[Pos{i, row}] = '#'
                    }
                }
            }

            if row > limit {
                limit = row
            }
            prev = Pos{col, row}
        }
    }

    return &Simulation{
        yLimit: limit,
        caveMap: caveMap,
        source: Pos{500, 0},
        sandUnitsAtRest: 0,
        currSandPos: Pos{500, 0},
    }
}

func Part_one(input_file string) int {
    Sim := initSimFromInput(input_file)

    for !Sim.Step(1) {
        continue
    }

    return Sim.sandUnitsAtRest
}

func Part_two(input_file string) int {
    Sim := initSimFromInput(input_file)

    for !Sim.Step(2) {
        continue
    }

    return Sim.sandUnitsAtRest
}

func main() {
    input := os.Args[1]
    fmt.Println(Part_one(input))
    fmt.Println(Part_two(input))
}

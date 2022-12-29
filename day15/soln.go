package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
    "sort"
)

type Pos [2]int // x, y

func AbsDiff(x, y int) int {
    if x > y {
        return x - y
    }
    return y - x
}

func Manhattan(p, q Pos) int {
    return AbsDiff(p[0], q[0]) + AbsDiff(p[1], q[1])
}

func mapImpossibles(sbmap map[Pos]Pos, atY int) map[int]bool {
    xVals := map[int]bool{}
    for s, b := range sbmap {
        radius := Manhattan(s, b)
        // We want to solve the inequality
        // Manhattan(s, (x, atY)) <= radius
        // for all x
        // rearranging...
        /// |s[0] - x| <= radius - |s[1] - atY|
        // -radius + |s[1] - atY| <= s[0] - x <= radius - |s[1] - atY|
        // radius - |s[1] - atY| + s[0] >= x >= -radius + |s[1] - atY| + s[0]
        //
        // this is the range of x values s.t. (x, atY) intersects the
        // circle between the sensor and its closest beacon, and thus
        // it cannot contain an unknown beacon (but can contain a known beacon)
        start := -radius + AbsDiff(s[1], atY) + s[0]
        end := radius - AbsDiff(s[1], atY) + s[0]

        for i := start; i <= end; i++ {
            xVals[i] = true
        }
    }
    return xVals
}


func countImpossibleBeacons(sbmap map[Pos]Pos, atY int) int {
    xVals := mapImpossibles(sbmap, atY)

    // get rid of spots that already have beacons in them
    for _, b := range sbmap {
        _, ok := xVals[b[0]]
        if b[1] == atY && ok {
            delete(xVals, b[0])
        }
    }

    return len(xVals)
}

func allContiguouslyOverlapping(sortedIntervals [][2]int) (bool, [2]int) {
    curr := sortedIntervals[0]
    for i := 1; i < len(sortedIntervals); i++ {
        next := sortedIntervals[i]
        if curr[1] >= next[0] && curr[1] >= next[1] {
            curr = curr
        } else if curr[1] >= next[0] {
            curr = [2]int{curr[0], next[1]}
        } else {
            return false, [2]int{curr[1], next[0]}
        }
    }
    return true, [2]int{}
}

// gauranteed exactly one possible position for the unknown beacon
func searchForUnknownBeacon(sbmap map[Pos]Pos, xLimit, yLimit int) Pos {
    for y := 0; y <= yLimit; y++ {
        ranges := [][2]int{}
        for s, b := range sbmap {
            radius := Manhattan(s, b)
            start := -radius + AbsDiff(s[1], y) + s[0]
            end := radius - AbsDiff(s[1], y) + s[0]
            if end < start {
                continue
            }
            if start < 0 {
                start = 0
            }
            if end > xLimit {
                end = xLimit
            }

            ranges = append(ranges, [2]int{start, end})
        }
        sort.Slice(ranges, func(i, j int) bool {
            return ranges[i][0] < ranges[j][0]
        })

        // gauranteed exactly one position not touched by radii
        // so check all x ranges for each y value until there is one
        // that doesn't cover the whole row
        if ok, val := allContiguouslyOverlapping(ranges); !ok {
            return Pos{val[0] + 1, y}
        }
    }

    return Pos{-1, -1}
}

func parseInput(input_file string) map[Pos]Pos {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    sensorClosestBeacon := map[Pos]Pos{}
    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()
        split := strings.Split(line, ": ")
        sensor := split[0][10:]
        beacon := split[1][21:]

        parsedCoord := [2]Pos{}
        for i, str := range [2]string{sensor, beacon} {
            split = strings.Split(str, ", ")
            x, err := strconv.Atoi(split[0][2:])
            if err != nil {
                log.Fatal(err)
            }
            y, err := strconv.Atoi(split[1][2:])
            if err != nil {
                log.Fatal(err)
            }
            parsedCoord[i] = Pos{x, y}
        }
        sensorClosestBeacon[parsedCoord[0]] = parsedCoord[1]
    }
    return sensorClosestBeacon
}

func Part_one(input_file string) int {
    sbmap := parseInput(input_file)
    return countImpossibleBeacons(sbmap, 2000000)
}

func Part_two(input_file string) int {
    sbmap := parseInput(input_file)
    c := searchForUnknownBeacon(sbmap, 4000000, 4000000)
    return 4000000*c[0] + c[1]
}


func main() {
    input := os.Args[1]
    fmt.Println(Part_one(input))
    fmt.Println(Part_two(input))
}

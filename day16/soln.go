package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

type Graph struct {
    keys map[string]int
    rates []int
    adjList [][]string
}

func copyMap(m map[string]bool) map[string]bool {
    m2 := make(map[string]bool, len(m))
    for k, v := range m {
        m2[k] = v
    }
    return m2
}


func recurseGraphWalk(graph *Graph, startKey string) int {
    var helper func(string, int, map[string]bool, int) int

    helper = func(
        currKey string,
        currMin int,
        released map[string]bool,
        totPressure int,
    ) int {
        if currMin == 0 {
            return totPressure
        }
        var r int
        turnedOff := false
        if _, ok := released[currKey]; !ok {
            idx := graph.keys[currKey]
            r = (currMin-1)*graph.rates[idx]
            turnedOff = true
        }

        notReleasedNew := copyMap(released)
        releasedNew := copyMap(released)
        releasedNew[currKey] = true

        best := 0
        idx := graph.keys[currKey]
        for _, n := range graph.adjList[idx] {
            res := helper(n, currMin-1, notReleasedNew, totPressure)
            if res > best {
                best = res
            }
            if turnedOff {
                res = helper(n, currMin-2, releasedNew, totPressure+r)
                if res > best {
                    best = res
                }
            }
        }
        return best
    }

    return helper(startKey, 30, map[string]bool{}, 0)
}

func parseInput(input_file string) *Graph {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    valveKeys := map[string]int{}
    valveRates := []int{}
    valveTunnels := [][]string{}
    scanner := bufio.NewScanner(file)
    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()
        // code here
        key := line[6:8]
        r := strings.Split(line, "=")
        rate := strings.Split(r[1], ";")[0]
        rateInt, err := strconv.Atoi(rate)
        if err != nil {
            log.Fatal(err)
        }

        var valves []string
        if strings.Contains(line, "valves ") {
            valves = strings.Split(line, "valves ")
        } else {
            valves = strings.Split(line, "valve ")
        }

        valves = strings.Split(valves[1], ", ")


        valveKeys[key] = len(valveKeys)
        valveRates = append(valveRates, rateInt)
        valveTunnels = append(valveTunnels, valves)
    }

    return &Graph{
        valveKeys,
        valveRates,
        valveTunnels,
    }
}

func Part_one(input_file string) int {
    G := parseInput(input_file)
    return recurseGraphWalk(G, "AA")
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

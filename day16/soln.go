package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
    "time"
)

type Graph struct {
    Keys []string
    AdjMatrix [][]int
    Rates []int
    Apsp [][]int //all pairs shortest path
}

// given a node that has rate 0
func (g *Graph) collapse(nodeIdx int) {
    if g.Rates[nodeIdx] != 0 {
        log.Fatal(fmt.Errorf("%s node has nonzero rate"), g.Keys[nodeIdx])
    }
    destCostMap := map[int]int{}
    for dest, cost := range g.AdjMatrix[nodeIdx] {
        if cost > 0 {
            destCostMap[dest] = cost
        }
    }

    for fromNode, row := range g.AdjMatrix {
        toCost := g.AdjMatrix[fromNode][nodeIdx]
        if toCost > 0 {
            for dest, cost := range destCostMap {
                if row[dest] == 0 {
                    row[dest] = toCost + cost
                } else {
                    row[dest] = Min(toCost + cost, row[dest])
                }
            }
        g.AdjMatrix[fromNode][nodeIdx] = 0
        }
    }
}

func (g *Graph) CollapseAllNodesWithZeroRate() {
    g.PrintAdjMatrix()
    for i, f := range g.Rates {
        if f == 0 {
            g.collapse(i)
        }
    }
    g.PrintAdjMatrix()
}

func (g *Graph) PrintAdjMatrix() {
    for _, row := range g.AdjMatrix {
        fmt.Println(row)
    }
    fmt.Print("\n")
}

func (g *Graph) SetAPSP(apsp [][]int) {
    g.Apsp = apsp
}

func Min(x, y int) int{
    if x < y {
        return x
    }
    return y
}


func pos(slice []string, key string) int {
    for p, v := range slice {
        if (v == key) {
            return p
        }
    }
    return -1
}

func AdjMatrix(keys []string, adjList [][]string) [][]int {
    matrix := [][]int{}
    for i, _ := range keys {
        row := make([]int, len(keys))
        for _, neighbKey := range adjList[i] {
            row[pos(keys, neighbKey)] = 1
        }
        matrix = append(matrix, row)
    }
    return matrix
}


func copy2D(src [][]int) (dst [][]int) {
	dst = make([][]int, len(src))
	for k := 0; k < len(src); k++ {
		dst[k] = make([]int, len(src[k]))
		copy(dst[k], src[k])
	}
	return dst
}

func AllPairsShortestPaths(matrix [][]int) [][]int {
    // Impl from https://github.com/bpowers/floyd-warshal/blob/4e6ff1dd0d79d68caff49b5c50db46d6f5dc1d86/floyd-warshal.go
    // initialize the shortest paths with a copy of the adjacency list
    prev := copy2D(matrix)
    curr := copy2D(matrix)

    nVertices := len(matrix)
    for k := 0; k < nVertices; k++ {
        // order of i,j iteration makes big difference
        for i := 0; i < nVertices; i++ {
            for j := 0; j < nVertices; j++ {
                a := prev[i][k] + prev[k][j]
                b := prev[i][j]
                if a < b {
                    curr[i][j] = a
                } else {
                    curr[i][j] = b
                }
            }
        }
        prev, curr = curr, prev
    }
    return prev
}

func copyMap(m map[int]int) map[int]int {
    m2 := make(map[int]int, len(m))
    for k, v := range m {
        m2[k] = v
    }
    return m2
}

// upper bound for branch and bound
func bestSurpassableInRemainingTime(
    graph *Graph,
    released map[int]int,
    currKey int,
    bestSoFar,
    currTot,
    currMin int) bool {
        sum := currTot
        for i := 0; i < len(graph.Keys); i++ {
            if _, ok := released[i]; !ok {
                dist := graph.Apsp[currKey][i]
                sum += (currMin-1-dist)*graph.Rates[i]
            }
        }
        //fmt.Println(sum > bestSoFar, sum, bestSoFar)
        return sum > bestSoFar
    }

func recurseGraphWalk(graph *Graph, startKey int) int {
    var helper func(int, int, map[int]int, int)
    bestSoFar := 0
    helper = func(
        currKey int,
        currMin int,
        released map[int]int,
        totPressure int,
    ) {

        if currMin <= 0 {
            if totPressure > bestSoFar {
                bestSoFar = totPressure
            }
            //fmt.Println(totPressure, bestSoFar)
            return
        }

        if !bestSurpassableInRemainingTime(
            graph,
            released,
            currKey,
            bestSoFar,
            totPressure,
            currMin,
        ) {
            return
        }

        var r int
         _, ok := released[currKey];
        if !ok {
            r = (currMin-1)*graph.Rates[currKey]
        }

        turnedOff := !ok && r > 0

        releasedNew := copyMap(released)
        if turnedOff {
            releasedNew[currKey] = currMin
        }

        for i, n := range graph.AdjMatrix[currKey] {
            if n > 0 {
                if !turnedOff {
                    helper(i, currMin-n, released, totPressure)
                } else {
                    helper(i, currMin-1-n, releasedNew, totPressure+r)
                }
            }
        }
    }
    start := time.Now()
    helper(startKey, 30, map[int]int{}, 0)
    elapsed := time.Since(start)
    fmt.Printf("Elapsed time: %s\n", elapsed)
    return bestSoFar
}

func parseInput(input_file string) *Graph {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    valveKeys := []string{}
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


        valveKeys = append(valveKeys, key)
        valveRates = append(valveRates, rateInt)
        valveTunnels = append(valveTunnels, valves)
    }
    matrix := AdjMatrix(valveKeys, valveTunnels)

    graph := &Graph{
        Keys: valveKeys,
        AdjMatrix: matrix,
        Rates: valveRates,
    }
    graph.CollapseAllNodesWithZeroRate()
    apsp := AllPairsShortestPaths(graph.AdjMatrix)
    graph.SetAPSP(apsp)
    return graph
}

func Part_one(input_file string) int {
    G := parseInput(input_file)

    return recurseGraphWalk(G, pos(G.Keys, "AA"))
}

func Part_two(input_file string) int {
    G := parseInput(input_file)
    return 0
}

func main() {
    input := os.Args[1]
    //fmt.Println(Part_one(input))
    fmt.Println(Part_two(input))

}

package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

type Crates struct {
    crates [][]string
}

func (c *Crates) addCrates(line string)  {
    // add more space if more is detected in the input
    totalInputBuckets := ((len(line) - 3) / 4) + 1
    existingBuckets := len(c.crates)
    newBuckets := totalInputBuckets - existingBuckets
    for i := 0; i < newBuckets; i++ {
        c.crates = append(c.crates, make([]string, 0))
    }
    // add crates to 2d array
    for pos, char := range line {
        if pos % 4 == 0 && char == '[' {
            idx := (pos / 4)
            c.crates[idx] = append(c.crates[idx], string(line[pos + 1]))
        }
    }
}

func (c *Crates) simulateMove(line string, part int) {
    res := strings.ReplaceAll(line, "move ", "")
    res = strings.ReplaceAll(res, "from ", "")
    res = strings.ReplaceAll(res, "to ", "")

    arr := strings.Split(res, " ")

    move, err := strconv.Atoi(arr[0])
    if err != nil {
        log.Fatal(err)
    }

    from, err := strconv.Atoi(arr[1])
    if err != nil {
        log.Fatal(err)
    }

    to, err := strconv.Atoi(arr[2])
    if err != nil {
        log.Fatal(err)
    }

    // We are 0 indexed, input is 1 indexed
    from -= 1
    to -= 1

    s := make([]string, move)
    copy(s, c.crates[from][0:move])
    c.crates[from] = c.crates[from][move:]

    if part == 1 {
        // reverse
        for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
            s[i], s[j] = s[j], s[i]
        }
    }

    c.crates[to] = append(s, c.crates[to]...)
}

func (c *Crates) topOfEachStack() string {
    s := ""
    for _, crate := range c.crates {
        if len(crate) > 0 {
            s = s + crate[0]
        }
    }
    return s
}

func Part_one(input_file string) string {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    simFlag := false
    crates := Crates{crates: make([][]string, 0)}
    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()

        // Check before feeding into Crate object
        if !simFlag && !strings.Contains(line, "[") {
            simFlag = true
            continue
        }
        if line == "" {
            continue
        }

        if simFlag {
            crates.simulateMove(line, 1)
        } else {
            crates.addCrates(line)
        }

    }
    return crates.topOfEachStack()
}

func Part_two(input_file string) string {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    simFlag := false
    crates := Crates{crates: make([][]string, 0)}
    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()

        // Check before feeding into Crate object
        if !simFlag && !strings.Contains(line, "[") {
            simFlag = true
            continue
        }
        if line == "" {
            continue
        }

        if simFlag {
            crates.simulateMove(line, 2)
        } else {
            crates.addCrates(line)
        }

    }
    return crates.topOfEachStack()
}

func main() {
    input := os.Args[1]
    fmt.Println(Part_one(input))
    fmt.Println(Part_two(input))
}

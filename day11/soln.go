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

type Monkey struct {
    Items []int
    Operation string // either "+" or "*"
    Operand int // if -1, then same as item
    TestMod int
    TrueTarget *Monkey
    FalseTarget *Monkey
    InspectedCount int
}

func (m *Monkey) TakeTurn(part int, modProduct int) {
    // Inspect each item
    for _, item := range m.Items {
        operand := m.Operand
        if operand == -1 {
            operand = item
        }

        worry := item
        if m.Operation == "+" {
            worry += operand
        } else {
            worry *= operand
        }

        if part == 1 {
            worry /= 3
        } else {
            worry %= modProduct
        }

        if worry % m.TestMod == 0 {
            m.ThrowTo(m.TrueTarget, worry)
        } else {
            m.ThrowTo(m.FalseTarget, worry)
        }

        m.InspectedCount++
    }
    m.Items = []int{}
}

func (m *Monkey) ThrowTo(reciever *Monkey, item int) {
    reciever.Items = append(reciever.Items, item)
}

func Solution(input_file string, part int) int {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    Monkeys := []*Monkey{&Monkey{}}

    parseCounter := -1
    trueTargets := []int{}
    falseTargets := []int{}
    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()
        currMonkey := Monkeys[len(Monkeys) - 1]
        if line == "end" {
            parseCounter = -1
            Monkeys = append(Monkeys, &Monkey{})
        } else {
            inp := strings.Split(line, ": ")
            switch parseCounter {
            case 0:
                if len(inp) == 1 {
                    currMonkey.Items = []int{}
                } else {
                    items := strings.Split(inp[1], ", ")
                    intItems := []int{}
                    for _, item := range items {
                        n, err := strconv.Atoi(item)
                        if err != nil {
                            log.Fatal(err)
                        }
                        intItems = append(intItems, n)
                    }
                    currMonkey.Items = intItems
                }
            case 1:
                fn := strings.Split(inp[1], "= ")
                var val string
                if strings.Contains(inp[1], "*") {
                    val = strings.Split(fn[1], " * ")[1]
                    currMonkey.Operation = "*"
                } else {
                    val = strings.Split(fn[1], " + ")[1]
                    currMonkey.Operation = "+"
                }
                if val == "old" {
                    currMonkey.Operand = -1
                } else {
                    vInt, err := strconv.Atoi(val)
                    if err != nil {
                        log.Fatal(err)
                    }
                    currMonkey.Operand = vInt
                }
            case 2:
                mod := strings.Split(inp[1], " by ")[1]
                modInt, err := strconv.Atoi(mod)
                if err != nil {
                    log.Fatal(err)
                }
                currMonkey.TestMod = modInt
            case 3:
                trueMonk := strings.Split(inp[1], "monkey ")[1]
                trueMonkInt, err := strconv.Atoi(trueMonk)
                if err != nil {
                    log.Fatal(err)
                }
                trueTargets = append(trueTargets, trueMonkInt)
            case 4:
                falseMonk := strings.Split(inp[1], "monkey ")[1]
                falseMonkInt, err := strconv.Atoi(falseMonk)
                if err != nil {
                    log.Fatal(err)
                }
                falseTargets = append(falseTargets, falseMonkInt)
            }
            parseCounter++
        }
    }

    // set the targets for throwing
    for i := 0; i < len(trueTargets); i++ {
        trueTarget := Monkeys[trueTargets[i]]
        falseTarget := Monkeys[falseTargets[i]]
        Monkeys[i].TrueTarget = trueTarget
        Monkeys[i].FalseTarget = falseTarget
    }

    // Part 1: 20, Part 2 : 10000
    var Rounds int
    if part == 1 {
        Rounds =  20
    } else {
        Rounds = 10000
    }
    // exploit the Chinese remainder theorem
    // so we don't have to keep track of overflowing worry levels.
    // The fact that all the moduli are coprime to each other is
    // significant because then the conditions for the CRT are satisfied
    // (I needed a hint for this)
    // A more verbose but less tricky solution is to keep track of the
    // remainder for each worry, for each Monkey's modulus, at each step
    // in the simulation, and to just check their respective remainders
    // at each Monkey's turn
    modProduct := 1
    for _, m := range Monkeys {
        modProduct *= m.TestMod
    }
    // simulate
    for r := 1; r <= Rounds; r++ {
        for _, m := range Monkeys {
            m.TakeTurn(part, modProduct)
        }
    }

    // compute monkey business
    inspections := []int{}
    for _, m := range Monkeys {
        inspections = append(inspections, m.InspectedCount)
    }
    sort.Ints(inspections)
    monkeyBusiness := inspections[len(inspections) - 1]*inspections[len(inspections) - 2]
    return monkeyBusiness
}


func main() {
    input := os.Args[1]
    // pt 1
    fmt.Println(Solution(input, 1))
    // pt 2
    fmt.Println(Solution(input, 2))
}

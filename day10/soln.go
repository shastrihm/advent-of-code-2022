package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

type Packet struct {
    TTL int
    ToAdd int
}

type CPUState struct {
    PacketQueue []*Packet
    Cycle int
    X int
    sigStrengthTotal int
}

type Coord struct {
    Row int
    Col int
}

type CRTState struct {
    CPU *CPUState
    CRT []string // convert this to 6 x 40 grid after rendering
    SpriteCenter *Coord
    CurrentPos *Coord
}

func (c *Coord) To1DCoord() int {
    return c.Row*40 + c.Col
}

func (c *Coord) From1DCoord(coord int) {
    c.Row = coord / 40
    c.Col = coord % 40
}

func (c *CPUState) Enqueue(pack *Packet) {
    c.PacketQueue = append(c.PacketQueue, pack)
}


func (c *CPUState) IncrementCycle() {
    c.Cycle += 1
}

func (c *CPUState) ComputeCycle() {
    // do signal strength check BEFORE going through the queue,
    // since the adds technically happen AFTER 2 cycles from addx

    if (c.Cycle + 20) % 40 == 0 && c.Cycle <= 220 {
         c.sigStrengthTotal += c.Cycle*c.X
    }

    // note that queue will be sorted by increasing TTL
    for i, p := range c.PacketQueue {
        if p.TTL == c.Cycle {
            c.X += p.ToAdd
        } else if p.TTL > c.Cycle {
            if i == len(c.PacketQueue) - 1 {
                c.PacketQueue = []*Packet{}
            } else {
                c.PacketQueue = c.PacketQueue[i:len(c.PacketQueue)]
            }
            break
        }
    }
}

func (c *CRTState) DrawNextPixel() {
    sprite := c.SpriteCenter.To1DCoord()
    current := c.CurrentPos.To1DCoord()
    if current >= len(c.CRT) {
        return
    }
    if c.CurrentPos.Col == sprite ||
        c.CurrentPos.Col == sprite + 1 ||
        c.CurrentPos.Col == sprite - 1 {
            c.CRT[current] = "#"
        } else {
            c.CRT[current] = "."
        }
    c.CurrentPos.From1DCoord(current + 1)
}

func (c *CRTState) UpdateSpritePos() {
    c.SpriteCenter.Col = c.CPU.X
}

func (c *CRTState) PrettyPrintCRT() {
    // Print
    for i, ch := range c.CRT {
        fmt.Print(string(ch), "")
        if (i + 1) % 40 == 0 {
            fmt.Print("\n")
        }
    }
    fmt.Print("\n")
}

func Part_one(input_file string) int {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    State := &CPUState{
        PacketQueue: []*Packet{},
        Cycle: 1,
        X: 1,
        sigStrengthTotal: 0,
    }
    readInputCycle := 0
    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()
        input := strings.Split(line, " ")
        instr := input[0]

        if instr == "noop" {
            readInputCycle += 1
        } else {
            // then instr = addx
            V, err := strconv.Atoi(input[1])
            if err != nil {
                log.Fatal(err)
            }
            State.Enqueue(&Packet{TTL: readInputCycle + 2, ToAdd: V})
            readInputCycle += 2
        }

    }

    for len(State.PacketQueue) > 0 {
        State.ComputeCycle()
        State.IncrementCycle()
    }
    return State.sigStrengthTotal
}


func Part_two(input_file string) []string {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    cpu := &CPUState{
        PacketQueue: []*Packet{},
        Cycle: 1,
        X: 1,
    }

    crt := &CRTState{
        CPU: cpu,
        CRT: make([]string, 40*6),
        SpriteCenter: &Coord{0, 1},
        CurrentPos: &Coord{0, 0},
    }

    readInputCycle := 0
    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()
        input := strings.Split(line, " ")
        instr := input[0]

        if instr == "noop" {
            readInputCycle += 1
        } else {
            // instr = addx
            V, err := strconv.Atoi(input[1])
            if err != nil {
                log.Fatal(err)
            }
            crt.CPU.Enqueue(&Packet{TTL: readInputCycle + 2, ToAdd: V})
            readInputCycle += 2
        }

    }

    for len(crt.CPU.PacketQueue) > 0 {
        crt.CPU.ComputeCycle()
        crt.DrawNextPixel()
        crt.CPU.IncrementCycle()
        crt.UpdateSpritePos()
    }

    // do the remaining
    for crt.CurrentPos.To1DCoord() < (40*6) {
        crt.DrawNextPixel()
        crt.CPU.IncrementCycle()
    }

    crt.PrettyPrintCRT()

    return crt.CRT
}

func main() {
    input := os.Args[1]
    fmt.Println(Part_one(input))
    fmt.Println(Part_two(input))
}

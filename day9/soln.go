package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
    "math"
)

type Coord struct {
    X float64
    Y float64
    PrevX float64
    PrevY float64
}

func (c *Coord) stringify() string {
    return fmt.Sprintf("%d %d", int(c.X), int(c.Y))
}

func sign(t float64) float64 {
    if t > 0 {
        return 1
    }
    return -1
}
// Given coords of head, tells tail what its next coordinates should be
func (tail *Coord) NextStep(head *Coord) {
    xdiff := sign(head.X - tail.X)
    ydiff := sign(head.Y - tail.Y)
    if math.Abs(head.X - tail.X) == 2 {
        tail.X += 1*xdiff
        if math.Abs(head.Y - tail.Y) == 1 {
            tail.Y += 1*ydiff
        }
    }
    if math.Abs(head.Y - tail.Y) == 2 {
        tail.Y += 1*ydiff
        if math.Abs(head.X - tail.X) == 1 {
            tail.X += 1*xdiff
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

    Head := &Coord{X:0, Y:0}
    Tail := &Coord{X:0, Y:0}

    TailVisited := map[string]bool{"0 0": true}
    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()
        arr := strings.Split(line, " ")
        dir := arr[0]

        step, err := strconv.Atoi(arr[1])
        if err != nil {
            log.Fatal(err)
        }
        for i := 0; i < step; i++ {
            Head.PrevX = Head.X
            Head.PrevY = Head.Y
            switch dir {
            case "R":
                Head.X += 1
            case "L":
                Head.X -= 1
            case "U":
                Head.Y += 1
            case "D":
                Head.Y -= 1
            }

            if math.Abs(Head.X - Tail.X) == 2 ||
                math.Abs(Head.Y - Tail.Y) == 2 {
                Tail.X = Head.PrevX
                Tail.Y = Head.PrevY
                stringTail := Tail.stringify()
                _, exists := TailVisited[stringTail]
                if !exists {
                    TailVisited[stringTail] = true
                }
            }
        }
    }
    return len(TailVisited)
}

func Part_two(input_file string) int {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var Rope []*Coord
    for i := 0; i < 10; i++ {
        Rope = append(Rope, &Coord{X:0, Y:0})
    }

    TailVisited := map[string]bool{"0 0": true}
    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()
        arr := strings.Split(line, " ")
        dir := arr[0]

        step, err := strconv.Atoi(arr[1])
        if err != nil {
            log.Fatal(err)
        }
        for i := 0; i < step; i++ {
            Head := Rope[0]
            Tail := Rope[1]
            Head.PrevX = Head.X
            Head.PrevY = Head.Y
            switch dir {
            case "R":
                Head.X += 1
            case "L":
                Head.X -= 1
            case "U":
                Head.Y += 1
            case "D":
                Head.Y -= 1
            }
            for seg := 0; seg < len(Rope) - 1; seg++  {
                Head = Rope[seg]
                Tail = Rope[seg + 1]
                Tail.NextStep(Head)
                if seg + 1 == 9 {
                    stringTail := Tail.stringify()
                    _, exists := TailVisited[stringTail]
                    if !exists {
                        TailVisited[stringTail] = true
                    }
                }
            }
        }
    }
    return len(TailVisited)
}


func main() {
    input := os.Args[1]
    fmt.Println(Part_one(input))
    fmt.Println(Part_two(input))
}

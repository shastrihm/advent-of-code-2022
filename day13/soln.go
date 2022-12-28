package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "encoding/json"
    "reflect"
    "sort"
)

func IsSlice(v interface{}) bool {
    return reflect.TypeOf(v).Kind() == reflect.Slice
}

func IsFloat(v interface{}) bool {
    return !IsSlice(v)
}
// 0 = wrong order
// 1 = right order
// 2 = keep going (recursion). also interpreted as equality
func ComparePair(arrL, arrR []interface{}) int {
    for i := 0; i < len(arrL); i++ {
        if i == len(arrR) {
            return 0
        }
        if IsFloat(arrL[i]) && IsFloat(arrR[i]) {
            if arrL[i].(float64) < arrR[i].(float64) {
                return 1
            } else if arrL[i].(float64) > arrR[i].(float64) {
                return 0
            }
        } else if IsSlice(arrL[i]) && IsSlice(arrR[i]) {
            res := ComparePair(arrL[i].([]interface{}), arrR[i].([]interface{}))
            if res != 2 {
                return res
            }
        } else {
            if IsFloat(arrL[i]) {
                res := ComparePair([]interface{}{arrL[i]}, arrR[i].([]interface{}))
                if res != 2 {
                    return res
                }
            } else {
                res := ComparePair(arrL[i].([]interface{}), []interface{}{arrR[i]})
                if res != 2 {
                    return res
                }
            }
        }
    }
    if len(arrR) > len(arrL) {
        return 1
    } else if len(arrR) == len(arrL) {
        return 2
    }
    return 0
}

func Part_one(input_file string) int {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    Pairs := [][]interface{}{}
    scanner := bufio.NewScanner(file)

    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()
        if line != "" {
            array := []interface{}{}
            err := json.Unmarshal([]byte(line), &array)
            if err != nil {
                log.Fatal(err)
            }
            Pairs = append(Pairs, array)
        }
    }
    s := 0
    for i := 0; i < len(Pairs); i += 2 {
        val := ComparePair(Pairs[i], Pairs[i+1])
        if val == 1 {
            s += (i + 2)/2
        }
    }
    return s
}

func Part_two(input_file string) int {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    Packets := [][]interface{}{}
    scanner := bufio.NewScanner(file)
    // Read input line by line
    for scanner.Scan() {
        line := scanner.Text()
        if line != "" {
            array := []interface{}{}
            err := json.Unmarshal([]byte(line), &array)
            if err != nil {
                log.Fatal(err)
            }
            Packets = append(Packets, array)
        }
    }

    idxs := []int{}
    for i := 0; i < len(Packets); i++ {
        idxs = append(idxs, i)
    }

    sort.Slice(idxs, func(i, j int) bool {
        res := ComparePair(Packets[idxs[i]], Packets[idxs[j]])
        if res == 0 {
            return false
        }
        return true
    })

    decoderKey := 1
    for i, idx := range idxs {
        res := fmt.Sprintf("%v", Packets[idx])
        if res == `[[6]]` || res == `[[2]]` {
            decoderKey *= i+1
        }
    }
    return decoderKey
}

func main() {
    input := os.Args[1]
    fmt.Println(Part_one(input))
    fmt.Println(Part_two(input))
}

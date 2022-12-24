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

type TreeNode struct {
    parent *TreeNode
    children []*TreeNode
    size int
    key string
}

func (node *TreeNode) AddChild(child *TreeNode) {
    node.children = append(node.children, child)
}

func (node *TreeNode) LookupChild(key string) *TreeNode {
    for _, child := range node.children {
        if child.key == key {
            return child
        }
    }
    return nil
}

func NewTreeNode(parent *TreeNode, size int, key string) *TreeNode {
    return &TreeNode{
        parent: parent,
        children: make([]*TreeNode, 0),
        size: size,
        key: key,
    }
}
// returns total size of all directories under 100000 (int) (for part 1)
// as well as a pointer to the root of the entire input filesystem to use
//      as input for part 2
func Part_one(input_file string) (int, *TreeNode) {
    file, err := os.Open(input_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // Read input line by line
    total := 0
    var current *TreeNode
    current = nil
    // input is just doing a depth-first search on the filesystem
    for scanner.Scan() {
        line := scanner.Text()
        if string(line[0]) == "$" && string(line[1:4]) == " cd" {
            nextDir := line[5:]
            if nextDir == "/" {
                current = NewTreeNode(nil, 0, "/")
            } else if nextDir != ".." {
                current = current.LookupChild(nextDir)
            } else {
                if current.parent != nil {
                    current.parent.size += current.size
                    if current.size <= 100000 {
                        total += current.size
                    }
                }
                current = current.parent
            }
        } else if string(line[0]) == "$" && string(line[1:4]) == " ls" {

        } else {
            // output of ls command
            var size int
            splitted := strings.Split(line, " ")
            key := splitted[1]
            if splitted[0] == "dir" {
                size = 0 // will increase during recursion
            } else {
                size, err = strconv.Atoi(splitted[0])
                if err != nil {
                    log.Fatal(err)
                }
            }
            node := NewTreeNode(current, size, key)
            current.AddChild(node)
            current.size += size
        }
    }
    // Do a final bubble-up from current
    var prev *TreeNode
    for current != nil {
        if current.parent != nil {
            current.parent.size += current.size
            if current.size <= 100000 {
                total += current.size
            }
        }
        prev = current
        current = current.parent
    }
    return total, prev
}

func dfs_helper(fs *TreeNode, currentDisk int, bestSoFar float64) float64 {
    if len(fs.children) == 0 {
        return bestSoFar
    }

    totalDisk := 70000000
    needDisk := 30000000
    diskSizeAfterPwdDeleted := totalDisk - (currentDisk - fs.size)
    if diskSizeAfterPwdDeleted >= needDisk && bestSoFar > float64(fs.size) {
        bestSoFar = float64(fs.size)
    }

    for _, child := range fs.children {
        bestSoFar = math.Min(bestSoFar, dfs_helper(child, currentDisk, bestSoFar))
    }

    return bestSoFar
}

func Part_two(filetree *TreeNode) float64 {
    currentDisk := filetree.size
    best := math.Inf(1)

    return dfs_helper(filetree, currentDisk, best)
}

func main() {
    input := os.Args[1]
    total, fs := Part_one(input)
    fmt.Println(total)
    fmt.Println(Part_two(fs))
}

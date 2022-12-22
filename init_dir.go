package main

import (
    "fmt"
    "log"
    "io"
    "os"
    "path/filepath"
)

// Given a day n
// Makes a new directory "day"+n
// With a readme, go file, and
// an empty input.txt and test_input.txt that
// you copy and paste into from the website
func main() {
    n := os.Args[1]

    // check if dir already exists
    if _, err := os.Stat("day"+n); !os.IsNotExist(err) {
        fmt.Printf("%s dir already exists, exiting...\n", "day"+n)
        os.Exit(3)
    }

    // make directory
    if err := os.Mkdir("day"+n, os.ModePerm); err != nil {
        log.Fatal(err)
    }

    // create readme
    fpath := filepath.Join("day"+n, "README.md")
    body := fmt.Sprintf("# Day %s \n\nProblem statement for day %s : https://adventofcode.com/2022/day/%s \n\n To run, cd into this directory and\n\n`go run . <input file>`\n\nwhere `<input file>` is either `input.txt` or `test_input.txt`", n, n, n)
    err := os.WriteFile(fpath, []byte(body), 0755)
    if err != nil {
        fmt.Printf("Unable to write file: %v", err)
    }

    // create soln.go file from templatae
     srcFile, err := os.Open("soln_template.go")
     if err != nil {
         log.Fatal(err)
     }
     defer srcFile.Close()

     fpath = filepath.Join("day"+n, "soln.go")
     destFile, err := os.Create(fpath)
     if err != nil {
         log.Fatal(err)
     }
     defer destFile.Close()

     _, err = io.Copy(destFile, srcFile)
     if err != nil {
         log.Fatal(err)
     }

     err = destFile.Sync()
     if err != nil {
         log.Fatal(err)
     }

     // create empty input.txt and test_input.txt
     // will have to copy and paste the input
     // from the website
     fpaths := []string{
         filepath.Join("day"+n, "input.txt"),
         filepath.Join("day"+n, "test_input.txt"),
     }
     for _, f := range fpaths {
         err := os.WriteFile(f, []byte("Copy and paste from https://adventofcode.com/2022/day/"+n), 0755)
         if err != nil {
             fmt.Printf("Unable to write file: %v", err)
         }
     }

     fmt.Println("Successfully created directory for day "+n)
}

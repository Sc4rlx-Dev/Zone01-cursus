package main

import (
    "flag"
    "fmt"
)

func init() {
    flag.Usage = func() {
        fmt.Println("Usage: go run . <input file> <output file>")
    }
}
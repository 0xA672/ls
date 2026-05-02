package main

import (
    "fmt"
    "os"
    "sort"
)

func main() {
    dir := "."
    if len(os.Args) > 1 {
        dir = os.Args[1]
    }

    entries, err := os.ReadDir(dir)
    if err != nil {
        fmt.Fprintf(os.Stderr, "ls: %v\n", err)
        os.Exit(1)
    }

    sort.Slice(entries, func(i, j int) bool {
        if entries[i].IsDir() != entries[j].IsDir() {
            return entries[i].IsDir()
        }
        return entries[i].Name() < entries[j].Name()
    })

    for _, e := range entries {
        name := e.Name()
        if e.IsDir() {
            name += "/"
        }
        fmt.Println(name)
    }
}

package main

import (
	"log"
	"os"
    "topasm/core"
)

func main() {
    if len(os.Args) < 2 { log.Fatal("Must provide a file to parse") }

    fileName := os.Args[1]

    file, err := os.ReadFile(fileName)
    if err != nil { log.Fatal(err) }

    fileLines := string(file)
    println(fileLines)

    token := core.Token {
        Type: core.Key,
        Str: "move",
        Position: uint64(0),
    }

    fault := core.NewFault(&token, "Test", "Nil")
    core.PrintFaults(fileName, []core.Fault { fault })
}

package main

import (
	"log"
	"os"
)

func main() {
    if len(os.Args) < 2 { log.Fatal("Must provide a file to parse") }

    fileName := os.Args[1]

    file, err := os.ReadFile(fileName)
    if err != nil { log.Fatal(err) }

    fileLines := string(file)
    println(fileLines)

    token := Token {
        Type: Key,
        Str: "move",
        Position: uint64(0),
    }

    fault := NewFault(&token, "Test", "Nil")
    PrintFaults(fileName, []Fault { fault })
}

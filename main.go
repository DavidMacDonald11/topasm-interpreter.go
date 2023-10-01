package main

import (
	"log"
	"os"
	"topasm/lexer"
	"topasm/util"
)

func main() {
    if len(os.Args) < 2 { log.Fatal("Must provide a file to parse") }
    filePath := os.Args[1]

    tokens, faults := lexer.TokenizeFile(filePath)

    println("Created tokens:")
    println(util.Join(tokens, ", ", "[", "]"))

    for _, fault := range faults {
        fault.Print()
    }
}

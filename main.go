package main

import (
	"log"
	"os"
	"topasm/core"
	"topasm/core/token"
	"topasm/lexer"
)

func main() {
    if len(os.Args) < 2 { log.Fatal("Must provide a file to parse") }
    fileName := os.Args[1]

    tokens, faults := lexer.TokenizeFile(fileName)

    println("Created tokens:")
    println(token.Join(tokens))

    core.PrintFaults(fileName, faults)
}

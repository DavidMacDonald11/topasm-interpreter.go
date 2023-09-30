package main

import (
	"log"
	"os"
	"topasm/lexer"
)

func main() {
    if len(os.Args) < 2 { log.Fatal("Must provide a file to parse") }
    filePath := os.Args[1]

    tokens, faults := lexer.TokenizeFile(filePath)

    println("Created tokens:")
    println(tokens.String())

    faults.Print(filePath)
}

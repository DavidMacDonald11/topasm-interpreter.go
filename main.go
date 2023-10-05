package main

import (
	"fmt"
	"log"
	"os"
	"topasm/interpreter"
	"topasm/lexer"
	"topasm/parser"
	"topasm/util"
)

func main() {
    if len(os.Args) < 2 { log.Fatal("Must provide a file to parse") }
    filePath := os.Args[1]

    tokens := lexer.TokenizeFile(filePath)

    println("Created tokens:")
    println(util.Join(tokens, ", ", "[", "]"))

    tree := parser.ParseTokens(tokens)

    println("Created Tree:")
    println(tree.String())

    fmt.Print("\nResult:\n")
    interpreter.InterpretTree(tree)
}

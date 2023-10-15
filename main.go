package main

import (
	"log"
	"os"
	"topasm/interpreter"
	"topasm/lexer"
	"topasm/parser"
)

func main() {
    if len(os.Args) < 2 { log.Fatal("Must provide a file to parse") }
    filePath := os.Args[1]

    tokens := lexer.TokenizeFile(filePath)
    tree := parser.ParseTokens(tokens)
    interpreter.InterpretTree(tree)
}

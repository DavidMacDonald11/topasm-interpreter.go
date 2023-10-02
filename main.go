package main

import (
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

    tokens, fault := lexer.TokenizeFile(filePath)

    println("Created tokens:")
    println(util.Join(tokens, ", ", "[", "]"))

    if fault != nil {
        fault.Print()
        return
    }

    tree, fault := parser.ParseTokens(tokens)

    println("Created Tree:")
    println(tree.String())

    if fault != nil {
        fault.Print()
        return
    }

    fault = interpreter.InterpretTree(tree)

    if fault != nil {
        fault.Print()
        return
    }
}

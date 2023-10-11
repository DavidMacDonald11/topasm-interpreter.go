package main

import (
	"flag"
	"log"
	"topasm/interpreter"
	"topasm/lexer"
	"topasm/parser"
)

func main() {
    filePath := flag.String("file", "", "the file to interpret")
    text := flag.String("text", "", "the text to interpret")
    flag.Parse()

    if *filePath == "" && *text == "" {
        log.Fatal("Must provide a file or text to parse")
    }

    var file lexer.SrcFile

    if *filePath != "" {
        file = lexer.OpenSrcFile(*filePath)
    } else {
        file = lexer.MakeSrcFile(*text)
    }

    tokens := lexer.TokenizeFile(&file)
    tree := parser.ParseTokens(tokens)
    interpreter.InterpretTree(tree)
}

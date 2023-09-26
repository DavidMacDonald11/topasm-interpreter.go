package lexer

import (
	"topasm/core"
	"topasm/core/grammar"
	"topasm/core/token"
	"topasm/io"
	"topasm/lexer/result"
)

type File = io.SourceFile
type Token = token.Token
type Fault = core.Fault
type Result = result.Result

func TokenizeFile(filePath string) ([]Token, []Fault) {
    file := io.NewSourceFile(filePath)
    tokens := []Token{}
    faults := []Fault{}

    for !file.AtEnd() {
        saveNextNewline := len(tokens) > 0 && !tokens[len(tokens) - 1].Has("\n")
        result := makeToken(file, saveNextNewline)

        if result.HasToken() { tokens = append(tokens, *result.Token) }
        if result.HasFault() { faults = append(faults, *result.Fault) }
        if result.Failed() { break }
    }

    tokens = append(tokens, newToken(file, token.Punc, grammar.EOF))
    return tokens, faults
}

func newToken(file File, kind token.Kind, str string) Token {
    return Token {
        Kind: kind,
        Str: str,
        Position: file.CharPos() - uint64(len(str)),
    }
}

func newFaultToken(file File, str string) Token {
    return newToken(file, token.None, str)
}

func makeToken(file File, saveNextNewline bool) Result {
    switch file.NextChar() {
    default:
        token := newFaultToken(file, file.ReadChar())
        return result.Failure(token, "Unrecognized symbol")
    }
}

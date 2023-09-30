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
type Tokens = token.Tokens
type Fault = core.Fault
type Faults = core.Faults
type Result = result.Result

func TokenizeFile(filePath string) (Tokens, Faults) {
    file := io.NewSourceFile(filePath)
    tokens := Tokens{}
    faults := Faults{}

    for !file.AtEnd() {
        saveNextNewline := len(tokens) > 0 && tokens.Last().Has("\n")
        result := makeToken(file, saveNextNewline)

        if result.HasToken() { tokens = append(tokens, *result.Token) }
        if result.HasFault() { faults = append(faults, *result.Fault) }
        if result.Failed() { break }
    }

    tokens = append(tokens, newToken(file, token.Punc, grammar.EOF))
    return tokens, faults
}

func newToken(file File, k token.Kind, s string) Token {
    return token.NewToken(k, s, file.CharPos() - uint64(len(s)))
}

func newFaultToken(file File, s string) Token {
    return newToken(file, token.None, s)
}

func makeToken(file File, saveNextNewline bool) Result {
    switch file.NextChar() {
    default:
        token := newFaultToken(file, file.ReadChar())
        return result.Failure(token, "Unrecognized symbol")
    }
}

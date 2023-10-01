package lexer

import (
	"strings"
	"topasm/fault"
	"topasm/grammar"
	"topasm/token"
	"topasm/util"
	"unicode"

	"golang.org/x/exp/slices"
)

type Token = token.Token
type Fault = fault.Fault

func TokenizeFile(path string) ([]Token, []Fault) {
    file := NewSrcFile(path)
    tokens := []Token{}
    faults := []Fault{}

    for !file.AtEnd() {
        saveNewline := len(tokens) > 0 && !tokens[len(tokens) - 1].Has("\n")
        result := makeToken(file, saveNewline)

        if result.HasToken() { tokens = append(tokens, *result.Token) }
        if result.HasFault() { faults = append(faults, *result.Fault) }
        if result.Failed() { break }
    }

    tokens = append(tokens, *token.New(token.Punc, grammar.EOF, file.Pos))
    return tokens, faults
}

func makeToken(file *SrcFile, saveNewline bool) *token.Result {
    for unicode.IsSpace(file.NextChar()) {
        if file.NextChar() == '\n' { break }
        file.ReadChar()
    }

    next := file.NextChar()

    switch {
    case next == ';':
        file.ReadCharsUntil("\n")
        return token.NoneResult()
    case next == '\n':
        return makeNewline(file, saveNewline)
    case strings.ContainsRune(grammar.Digits, next):
        return makeNum(file)
    case strings.ContainsRune(grammar.Letters, next):
        return makeIdOrKey(file)
    case strings.ContainsRune(grammar.Puncs, next):
        return makePunc(file)
    }

    tok := token.New(token.None, file.ReadChar(), file.Pos)
    return token.FailureResult(tok, "Lexing", "Unrecognized symbol")
}

func makeNewline(file *SrcFile, saveNewline bool) *token.Result {
    tok := token.New(token.Punc, file.ReadChar(), file.Pos)

    if !saveNewline { return token.NoneResult() }
    return token.TokenResult(tok)
}

func makeNum(file *SrcFile) *token.Result {
    str := file.ReadTheseChars(grammar.Digits)
    tok := token.New(token.Num, str, file.Pos)
    return token.TokenResult(tok)
}

func makeIdOrKey(file *SrcFile) *token.Result {
    str := file.ReadTheseChars(grammar.Letters)
    isKey := slices.Contains(grammar.Keys(), str)
    kind := util.IfElse(isKey, token.Key, token.Id)
    return token.TokenResult(token.New(kind, str, file.Pos))
}

func makePunc(file *SrcFile) *token.Result {
    tok := token.New(token.Punc, file.ReadChar(), file.Pos)
    return token.TokenResult(tok)
}

package lexer

import (
	"strings"
	"topasm/core"
	"topasm/core/grammar"
	"topasm/core/token"
	"topasm/io"
	"topasm/lexer/result"
	"unicode"
	"golang.org/x/exp/slices"
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
        result := makeToken(&file, saveNextNewline)

        if result.HasToken() { tokens = append(tokens, *result.Token) }
        if result.HasFault() { faults = append(faults, *result.Fault) }
        if result.Failed() { break }
    }

    tokens = append(tokens, newToken(&file, token.Punc, grammar.EOF))
    return tokens, faults
}

func newToken(file *File, k token.Kind, s string) Token {
    return token.NewToken(k, s, file.CharPos() - uint64(len(s)))
}

func newFaultToken(file *File, s string) Token {
    return newToken(file, token.None, s)
}

func makeToken(file *File, saveNextNewline bool) Result {
    for unicode.IsSpace(file.NextChar()) {
        if file.NextChar() == '\n' { break }
        file.ReadChar()
    }

    next := file.NextChar()

    switch {
    case next == ';':
        file.ReadCharsUntil("\n")
        return result.None()
    case next == '\n':
        token := newToken(file, token.Punc, file.ReadChar())
        return core.IfElse(saveNextNewline, result.Token(token), result.None())
    case strings.ContainsRune(grammar.Digits, next):
        return makeNum(file)
    case strings.ContainsRune(grammar.Letters, next):
        return makeIdOrKey(file)
    case strings.ContainsRune(grammar.Puncs, next):
        return makePunc(file)
    default:
        token := newFaultToken(file, file.ReadChar())
        return result.Failure(token, "Unrecognized symbol")
    }
}

func makeNum(file *File) Result {
    str := file.ReadTheseChars(grammar.Digits)
    token := newToken(file, token.Num, str)
    return result.Token(token)
}

func makeIdOrKey(file *File) Result {
    str := file.ReadTheseChars(grammar.Letters)
    isKeyword := slices.Contains(grammar.Keywords(), str)
    kind := core.IfElse(isKeyword, token.Key, token.Id)

    return result.Token(newToken(file, kind, str))
}

func makePunc(file *File) Result {
    token := newToken(file, token.Punc, file.ReadChar())
    return result.Token(token)
}

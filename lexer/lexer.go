package lexer

import (
	"strings"
	"topasm/grammar"
	"topasm/token"
	"topasm/util"
	"unicode"

	"golang.org/x/exp/slices"
)

func TokenizeFile(path string) []token.Token {
    file := NewSrcFile(path)
    tokens := []token.Token{}

    for {
        saveNewline := len(tokens) > 0 && !tokens[len(tokens) - 1].Has("\n")
        tok := makeToken(&file)

        if saveNewline || !tok.Has("\n") { tokens = append(tokens, tok) }
        if tok.Has(grammar.EOF) { return tokens }
    }
}

func makeToken(file *SrcFile) token.Token {
    next := file.NextChar()

    switch {
    case next == '\u0000':
        fallthrough
    case strings.ContainsRune(grammar.Puncs, next):
        return makePunc(file)
    case strings.ContainsRune(grammar.Digits, next):
        return makeNum(file)
    case strings.ContainsRune(grammar.Letters, next):
        return makeIdOrKey(file)
    case next == '\'':
        return makeChar(file)
    case next == ';':
        file.ReadCharsUntil("\n")
        return makeToken(file)
    case unicode.IsSpace(next):
        file.ReadChar()
        return makeToken(file)
    }

    tok := token.New(token.None, file.ReadChar(), file.Line)
    util.Fail(tok, "Unrecognized symbol")
    return tok
}

func makePunc(file *SrcFile) token.Token {
    line := file.Line
    str := util.IfElse(file.AtEnd(), grammar.EOF, file.ReadChar())
    return token.New(token.Punc, str, line)
}

func makeNum(file *SrcFile) token.Token {
    str := file.ReadTheseChars(grammar.Digits)
    return token.New(token.Num, str, file.Line)
}

func makeIdOrKey(file *SrcFile) token.Token {
    str := file.ReadTheseChars(grammar.Letters)
    isKey := slices.Contains(grammar.Keys(), str)
    kind := util.IfElse(isKey, token.Key, token.Id)
    return token.New(kind, str, file.Line)
}

func makeChar(file *SrcFile) token.Token {
    b := strings.Builder{}
    b.WriteString(file.ReadChar())

    esc := file.NextChar() == '\\'
    b.WriteString(file.ReadChars(util.IfElse(esc, 2, 1)))

    q := file.ReadChar()
    b.WriteString(q)

    tok := token.New(token.Char, b.String(), file.Line)
    if q != "'" { util.Fail(tok, "Unterminated char") }
    return tok
}

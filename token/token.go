package token

import (
	"fmt"
	"strings"
	"topasm/util"

	"golang.org/x/exp/slices"
)

type Kind string

const (
    Punc Kind = "P"
    Id Kind = "I"
    Key Kind = "K"
    Num Kind = "N"
    Char Kind = "C"
    None Kind = "?"
)

func (k Kind) String() string { return string(k) }

type Token struct {
    Kind Kind
    Str string
    Line int
}

func New(kind Kind, str string, pos int) Token {
    return Token{kind, str, pos}
}

func (t Token) Position() int {
    return t.Line
}

func (t Token) Of(kinds ...Kind) bool {
    return slices.Contains(kinds, t.Kind)
}

func (t Token) Has(strs ...string) bool {
    return slices.Contains(strs, t.Str)
}

func (t Token) String() string {
    escapedStr := strings.ReplaceAll(t.Str, "\n", "\\n")
    str := util.IfElse(t.Str == "", "EOF", escapedStr)
    return fmt.Sprintf("%s'%s'", t.Kind, str)
}

func (t Token) RecString(prefix string) string {
    return t.String()
}

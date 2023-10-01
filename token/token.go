package token

import (
	"fmt"
	"strings"
	"topasm/fault"
	"topasm/util"

	"golang.org/x/exp/slices"
)

type Kind string

const (
    Punc Kind = "P"
    Id Kind = "I"
    Key Kind = "K"
    Num Kind = "N"
    None Kind = "?"
)

type Token struct {
    Kind Kind
    Str string
    Pos int
}

func New(k Kind, s string, pos int) *Token {
    return &Token{Kind: k, Str: s, Pos: pos}
}

func (t *Token) Position() fault.Position {
    return *fault.NewPosition(t.Pos - len(t.Str) + 1, t.Pos)
}

func (t *Token) Of(kinds ...Kind) bool {
    return slices.Contains(kinds, t.Kind)
}

func (t *Token) Has(strs ...string) bool {
    return slices.Contains(strs, t.Str)
}

func (t Token) String() string {
    escapedStr := strings.ReplaceAll(t.Str, "\n", "\\n")
    str := util.IfElse(t.Str == "", "EOF", escapedStr)
    return fmt.Sprintf("%s'%s'", t.Kind, str)
}

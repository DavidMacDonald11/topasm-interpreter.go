package token

import (
	"fmt"
	"strings"
	"golang.org/x/exp/slices"
    "topasm/core"
)

type Kind string

const (
    Punc Kind = "P"
    Id Kind = "I"
    Key Kind = "K"
    Num Kind = "N"
    Char Kind = "C"
    Str Kind = "S"
    None Kind = "?"
)

type Token struct {
    kind Kind
    str string
    pos uint64
}

func NewToken(k Kind, s string, pos uint64) Token {
    return Token {
        kind: k,
        str: s,
        pos: pos,
    }
}

func (t *Token) Position() core.UIntRange {
    return core.UIntRange {
        Start: t.pos,
        End: t.pos + uint64(len(t.str) - 1),
    }
}

func (t *Token) Of(kinds ...Kind) bool {
    return slices.Contains(kinds, t.kind)
}

func (t *Token) Has(strs ...string) bool {
    return slices.Contains(strs, t.str)
}

func (t *Token) String() string {
    str := core.IfElse(t.str == "", "EOF", t.escapedStr())
    return fmt.Sprintf("%s'%s'", t.kind, str)
}

func (t *Token) escapedStr() string {
    return strings.ReplaceAll(t.str, "\n", "\\n")
}

type Tokens []Token

func (t Tokens) String() string {
    builder := strings.Builder{}

    for i, token := range t {
        if i != 0 { builder.WriteString(", ") }
        builder.WriteString(token.String())
    }

    return "[" + builder.String() + "]"
}

func (t Tokens) Last() *Token {
    return core.IfElse(len(t) > 0, &t[len(t) - 1], nil)
}

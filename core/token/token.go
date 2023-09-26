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
    Kind Kind
    Str string
    Position uint64
}

func Join(tokens []Token) string {
    builder := strings.Builder{}

    for i, token := range tokens {
        if i != 0 { builder.WriteString(", ") }
        builder.WriteString(token.String())
    }

    return "[" + builder.String() + "]"
}

func (self *Token) FaultPosition() core.UIntRange {
    return core.UIntRange {
        Start: self.Position,
        End: self.Position + uint64(len(self.Str) - 1),
    }
}

func (self *Token) Of(kinds ...Kind) bool {
    return slices.Contains(kinds, self.Kind)
}

func (self *Token) Has(strs ...string) bool {
    return slices.Contains(strs, self.Str)
}

func (self *Token) String() string {
    str := core.IfThen(self.Str == "", "EOF", self.escapedStr())
    return fmt.Sprintf("%s'%s'", self.Kind, str)
}

func (self *Token) escapedStr() string {
    return strings.ReplaceAll(self.Str, "\n", "\\n")
}

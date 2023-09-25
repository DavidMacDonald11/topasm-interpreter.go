package core

import (
	"fmt"
	"strings"
	"golang.org/x/exp/slices"
)

type TokenType string

const (
    Punc TokenType = "P"
    Id TokenType = "I"
    Key TokenType = "K"
    Num TokenType = "N"
    Char TokenType = "C"
    Str TokenType = "S"
    None TokenType = "?"
)

type Token struct {
    Type TokenType
    Str string
    Position uint64
}

func (self *Token) FaultPosition() UIntRange {
    return UIntRange {
        Start: self.Position,
        End: self.Position + uint64(len(self.Str) - 1),
    }
}

func (self *Token) Of(types ...TokenType) bool {
    return slices.Contains(types, self.Type)
}

func (self *Token) Has(strs ...string) bool {
    return slices.Contains(strs, self.Str)
}

func (self *Token) String() string {
    str := IfThen(self.Str == "", "EOF", self.escapedStr())
    return fmt.Sprintf("%s'%s'", self.Type, str)
}

func (self *Token) escapedStr() string {
    return strings.ReplaceAll(self.Str, "\n", "\\n")
}

package main

import (
	"fmt"
	"strings"
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
    for _, t := range types {
        if self.Type == t { return true }
    }

    return false
}

func (self *Token) Has(strs ...string) bool {
    for _, s := range strs {
        if self.Str == s { return true }
    }

    return false
}

func (self *Token) String() string {
    var str string

    if self.Str == "" {
        str = "EOF"
    } else {
        str = strings.ReplaceAll(self.Str, "\n", "\\n")
    }

    return fmt.Sprintf("%s'%s'", self.Type, str)
}

package parser

import (
	"fmt"
	"strings"
	"topasm/fault"
	"topasm/token"
	"topasm/util"
)

type Context struct {
    tokens []token.Token
    n int
}

func NewContext(tokens []token.Token) *Context {
    return &Context{tokens, 0}
}

func (c *Context) Next() *token.Token {
    return &c.tokens[c.n]
}

func (c *Context) Take() *token.Token {
    tok := c.Next()
    c.n += 1
    return tok
}

func (c *Context) ExpectingOf(kinds ...token.Kind) (*Token, *Fault) {
    tok := c.Take()
    if tok.Of(kinds...) { return tok, nil }

    list := util.Join(kinds, ", ", "[", "]")
    msg := fmt.Sprintf("Expecting one of kind %s", list)

    return tok, fault.New(tok, "Parsing", msg)
}

func (c *Context) ExpectingHas(strs ...string) (*Token, *Fault) {
    tok := c.Take()
    if tok.Has(strs...) { return tok, nil }

    strs = util.Map(strs, func(it string) string {
        return strings.ReplaceAll(it, "\n", "\\n")
    })

    list := util.JoinStr(strs, ", ", "[", "]")
    msg := fmt.Sprintf("Expecting one of %s", list)

    return tok, fault.New(tok, "Parsing", msg)
}

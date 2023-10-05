package parser

import (
	"fmt"
	"strings"
	"topasm/token"
	"topasm/util"
)

type Context struct {
    tokens []token.Token
    n int
}

func NewContext(tokens []token.Token) Context {
    return Context{tokens, 0}
}

func (c Context) Next() token.Token {
    return c.tokens[c.n]
}

func (c *Context) Take() token.Token {
    tok := c.Next()
    if c.n < len(c.tokens) - 1 { c.n += 1 }
    return tok
}

func (c *Context) ExpectingOf(kinds ...token.Kind) token.Token {
    tok := c.Take()
    if tok.Of(kinds...) { return tok }

    list := util.Join(kinds, ", ", "[", "]")
    msg := fmt.Sprintf("Expecting one of kind %s", list)

    util.Fail(tok, msg)
    return tok
}

func (c *Context) ExpectingHas(strs ...string) token.Token {
    tok := c.Take()
    if tok.Has(strs...) { return tok }

    strs = util.Map(strs, func(it string) string {
        return strings.ReplaceAll(it, "\n", "\\n")
    })

    list := util.JoinStr(strs, ", ", "[", "]")
    msg := fmt.Sprintf("Expecting one of %s", list)

    util.Fail(tok, msg)
    return tok
}

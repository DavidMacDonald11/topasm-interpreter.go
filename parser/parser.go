package parser

import (
	"topasm/fault"
	"topasm/node"
	"topasm/token"
)

type Fault = fault.Fault
type Token = token.Token
type Node = node.Node
type Builder = node.Builder

func ParseTokens(tokens []Token) (*Node, *Fault) {
    ctx := NewContext(tokens)
    r := parseMove(ctx)
    return r.Node, r.Fault
}

func parseMove(ctx *Context) *node.Result {
    b := node.NewBuilder("Move")

    b.Add(ctx.ExpectingHas("move"))
    b.AddResult(parseValue(ctx))
    b.Add(ctx.ExpectingHas("into"))
    b.AddResult(parseReg(ctx))

    return b.Result()
}

func parseValue(ctx *Context) *node.Result {
    if ctx.Next().Has("#") { return parseReg(ctx) }
    return parseNum(ctx)
}

func parseReg(ctx *Context) *node.Result {
    b := node.NewBuilder("Reg")

    b.Add(ctx.ExpectingHas("#"))
    b.Add(ctx.ExpectingOf(token.Num))

    return b.Result()
}

func parseNum(ctx *Context) *node.Result {
    b := node.NewBuilder("Num")
    b.Add(ctx.ExpectingOf(token.Num))

    return b.Result()
}

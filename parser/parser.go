package parser

import (
	"topasm/fault"
	"topasm/grammar"
	"topasm/node"
	"topasm/token"
)

type Fault = fault.Fault
type Token = token.Token
type Node = node.Node

func ParseTokens(tokens []Token) (*Node, *Fault) {
    ctx := NewContext(tokens)
    b := node.NewBuilder("File")

    for !ctx.Next().Has(grammar.EOF) {
        b.AddResult(parseIns(ctx))
        b.Add(ctx.ExpectingHas("\n", grammar.EOF))
    }

    r := b.Result()
    return r.Node, r.Fault
}

func parseIns(ctx *Context) *node.Result {
    switch {
    case ctx.Next().Has("move"): return parseMove(ctx)
    case ctx.Next().Has("add"): return parseAdd(ctx)
    case ctx.Next().Has("sub"): return parseSub(ctx)
    case ctx.Next().Has("printc"): return parsePrintc(ctx)
    case ctx.Next().Has("printi"): return parsePrinti(ctx)
    }

    b := node.NewBuilder("Bad Node")
    b.AddChild(ctx.Take())
    n := b.Result().Node

    return node.FaultResult(n, "Parsing", "Unexpected token")
}

func parseMove(ctx *Context) *node.Result {
    return parseBinary(ctx, "move", "into")
}

func parseAdd(ctx *Context) *node.Result {
    return parseBinary(ctx, "add", "into")
}

func parseSub(ctx *Context) *node.Result {
    return parseBinary(ctx, "sub", "from")
}

func parsePrintc(ctx *Context) *node.Result {
    return parseUnary(ctx, "printc")
}

func parsePrinti(ctx *Context) *node.Result {
    return parseUnary(ctx, "printi")
}

func parseUnary(ctx *Context, ins string) *node.Result {
    b := node.NewBuilder("Unary Instruction")

    b.Add(ctx.ExpectingHas(ins))
    b.AddResult(parseValue(ctx))

    return b.Result()
}

func parseBinary(ctx *Context, ins string, prep string) *node.Result {
    b := node.NewBuilder("Binary Instruction")

    b.Add(ctx.ExpectingHas(ins))
    b.AddResult(parseValue(ctx))
    b.Add(ctx.ExpectingHas(prep))
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

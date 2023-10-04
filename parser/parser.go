package parser

import (
	"topasm/fault"
	"topasm/grammar"
	"topasm/node"
	"topasm/token"
)

func ParseTokens(tokens []token.Token) node.File {
    ctx := NewContext(tokens)
    nodes := []node.Node{}

    for !ctx.Next().Has(grammar.EOF) {
        ins := parseIns(&ctx)
        lf := ctx.ExpectingHas("\n", grammar.EOF)
        nodes = append(nodes, ins, lf)
    }

    return node.NewFile(nodes)
}

func parseIns(ctx *Context) node.Node {
    switch {
    case ctx.Next().Has("move"):
        return parseBinary(ctx, "move", "into")
    case ctx.Next().Has("add"):
        return parseBinary(ctx, "add", "into")
    case ctx.Next().Has("sub"):
        return parseBinary(ctx, "sub", "from")
    case ctx.Next().Has("inc"):
        return parseUnary(ctx, "inc")
    case ctx.Next().Has("dec"):
        return parseUnary(ctx, "dec")
    case ctx.Next().Has("printc"):
        return parseUnary(ctx, "printc")
    case ctx.Next().Has("printi"):
        return parseUnary(ctx, "printi")
    }

    err := node.NewError(ctx.Take())
    fault.Fail(err, "Parsing", "Unexpected token")
    return err
}

func parseUnary(ctx *Context, ins string) node.UnaryIns {
    insTok := ctx.ExpectingHas(ins)
    reg := parseReg(ctx)

    return node.NewUnaryIns(insTok, reg)
}

func parseBinary(ctx *Context, ins string, prep string) node.BinaryIns {
    insTok := ctx.ExpectingHas(ins)
    val := parseValue(ctx)
    prepTok := ctx.ExpectingHas(prep)
    reg := parseReg(ctx)

    return node.NewBinaryIns(insTok, val, prepTok, reg)
}

func parseValue(ctx *Context) node.Node {
    if ctx.Next().Has("#") { return parseReg(ctx) }
    return parseNum(ctx)
}

func parseReg(ctx *Context) node.Reg {
    hash := ctx.ExpectingHas("#")
    num := ctx.ExpectingOf(token.Num)
    return node.NewReg(hash, num)
}

func parseNum(ctx *Context) node.Num {
    num := ctx.ExpectingOf(token.Num)
    return node.NewNum(num)
}

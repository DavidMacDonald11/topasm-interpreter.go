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
    case ctx.Next().Has("move"): return parseMove(ctx)
    case ctx.Next().Has("add"): return parseAdd(ctx)
    case ctx.Next().Has("sub"): return parseSub(ctx)
    case ctx.Next().Has("printc"): return parsePrintc(ctx)
    case ctx.Next().Has("printi"): return parsePrinti(ctx)
    }

    err := node.NewError(ctx.Take())
    fault.Fail(err, "Parsing", "Unexpected token")
    return err
}

func parseMove(ctx *Context) node.Node {
    return parseBinary(ctx, "move", "into")
}

func parseAdd(ctx *Context) node.Node {
    return parseBinary(ctx, "add", "into")
}

func parseSub(ctx *Context) node.Node {
    return parseBinary(ctx, "sub", "from")
}

func parsePrintc(ctx *Context) node.Node {
    return parseUnary(ctx, "printc")
}

func parsePrinti(ctx *Context) node.Node {
    return parseUnary(ctx, "printi")
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

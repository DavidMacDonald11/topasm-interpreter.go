package parser

import (
	"strings"
	"topasm/grammar"
	"topasm/node"
	"topasm/token"
	"topasm/util"
)

func ParseTokens(tokens []token.Token) node.Node {
    ctx := NewContext(tokens)
    nodes := []node.RecStringer{}

    for !ctx.Next().Has(grammar.EOF) {
        ins := parseIns(&ctx)
        lf := ctx.ExpectingHas("\n", grammar.EOF)
        nodes = append(nodes, ins, lf)
    }

    return node.New("File", nodes...)
}

func parseIns(ctx *Context) node.Node {
    switch {
    case ctx.Next().Of(token.Id):
        return parseLabel(ctx)
    case ctx.Next().Has("call"):
        return parseCall(ctx)
    case ctx.Next().Has("move"):
        return parseInsValReg(ctx, "Move", "into")
    case ctx.Next().Has("add"):
        return parseInsValReg(ctx, "Add", "into")
    case ctx.Next().Has("sub"):
        return parseInsValReg(ctx, "Sub", "from")
    case ctx.Next().Has("inc"):
        return parseInsReg(ctx, "Inc")
    case ctx.Next().Has("dec"):
        return parseInsReg(ctx, "Dec")
    }

    err := node.New("?", ctx.Take())
    util.Fail(err, "Unexpected token")
    return err
}

func parseLabel(ctx *Context) node.Node {
    id := ctx.ExpectingOf(token.Id)
    colon := ctx.ExpectingHas(":")
    return node.New("Label", id, colon)
}

func parseCall(ctx *Context) node.Node {
    ins := ctx.ExpectingHas("call")
    id := ctx.ExpectingOf(token.Id)
    return node.New("Call", ins, id)
}

func parseInsReg(ctx *Context, name string) node.Node {
    ins := ctx.ExpectingHas(strings.ToLower(name))
    reg := parseReg(ctx)
    return node.New(name, ins, reg)
}

func parseInsValReg(ctx *Context, name string, prep string) node.Node {
    ins := ctx.ExpectingHas(strings.ToLower(name))
    val := parseValue(ctx)
    p := ctx.ExpectingHas(prep)
    reg := parseReg(ctx)

    return node.New(name, ins, val, p, reg)
}

func parseValue(ctx *Context) node.Node {
    if ctx.Next().Has("#") { return parseReg(ctx) }
    return parseNum(ctx)
}

func parseReg(ctx *Context) node.Node {
    hash := ctx.ExpectingHas("#")
    num := ctx.ExpectingOf(token.Num)
    return node.New("Reg", hash, num)
}

func parseNum(ctx *Context) node.Node {
    num := ctx.ExpectingOf(token.Num)
    return node.New("Num", num)
}

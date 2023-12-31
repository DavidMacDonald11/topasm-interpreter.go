package parser

import (
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
    case ctx.Next().Has(grammar.JumpKeys()...):
        return parseJump(ctx)
    case ctx.Next().Has("call"):
        return parseCall(ctx)
    case ctx.Next().Has("move"):
        return parseInsValReg(ctx, "move", "into")
    case ctx.Next().Has("add"):
        return parseInsValReg(ctx, "add", "into")
    case ctx.Next().Has("sub"):
        return parseInsValReg(ctx, "sub", "from")
    case ctx.Next().Has("comp"):
        return parseInsValVal(ctx, "comp", "with")
    case ctx.Next().Has("inc"):
        return parseInsReg(ctx, "inc")
    case ctx.Next().Has("dec"):
        return parseInsReg(ctx, "dec")
    case ctx.Next().Has("mul"):
        return parseInsReg(ctx, "mul")
    case ctx.Next().Has("div"):
        return parseInsReg(ctx, "div")
    }

    err := node.New("?", ctx.Take())
    util.Fail(err, "Unexpected token")
    return err
}

func parseLabel(ctx *Context) node.Node {
    id := ctx.ExpectingOf(token.Id)
    colon := ctx.ExpectingHas(":")
    return node.New("label", id, colon)
}

func parseJump(ctx *Context) node.Node {
    ins := ctx.ExpectingHas(grammar.JumpKeys()...)
    id := ctx.ExpectingOf(token.Id)
    return node.New("jump", ins, id)
}

func parseCall(ctx *Context) node.Node {
    ins := ctx.ExpectingHas("call")
    id := ctx.ExpectingOf(token.Id)
    return node.New("call", ins, id)
}

func parseInsReg(ctx *Context, name string) node.Node {
    ins := ctx.ExpectingHas(name)
    reg := parseReg(ctx)
    return node.New(name, ins, reg)
}

func parseInsValReg(ctx *Context, name string, prep string) node.Node {
    ins := ctx.ExpectingHas(name)
    val := parseValue(ctx)
    p := ctx.ExpectingHas(prep)
    reg := parseReg(ctx)

    return node.New(name, ins, val, p, reg)
}

func parseInsValVal(ctx *Context, name string, prep string) node.Node {
    ins := ctx.ExpectingHas(name)
    v1 := parseValue(ctx)
    p := ctx.ExpectingHas(prep)
    v2 := parseValue(ctx)

    return node.New(name, ins, v1, p, v2)
}

func parseValue(ctx *Context) node.Node {
    if ctx.Next().Has("#") { return parseReg(ctx) }
    if ctx.Next().Of(token.Char) { return parseChar(ctx) }
    return parseNum(ctx)
}

func parseReg(ctx *Context) node.Node {
    hash := ctx.ExpectingHas("#")
    num := ctx.ExpectingOf(token.Num)
    return node.New("reg", hash, num)
}

func parseChar(ctx *Context) node.Node {
    char := ctx.ExpectingOf(token.Char)
    return node.New("char", char)
}

func parseNum(ctx *Context) node.Node {
    num := ctx.ExpectingOf(token.Num)
    return node.New("num", num)
}

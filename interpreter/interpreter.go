package interpreter

import (
	"fmt"
	"strconv"
	"topasm/node"
	"topasm/token"
	"topasm/util"
)

func InterpretTree(tree node.Node) {
    ctx := NewContext(tree)

    for ; ctx.i * 2 < len(tree.Children); ctx.i++ {
        ins := tree.Children[ctx.i * 2].(node.Node)

        switch ins.Name {
        case "Label": continue
        case "Call": intCall(&ctx, ins)
        case "Move": intMove(&ctx, ins)
        case "Add": intAdd(&ctx, ins)
        case "Sub": intSub(&ctx, ins)
        case "Inc": intInc(&ctx, ins)
        case "Dec": intDec(&ctx, ins)
        default: util.Fail(ins, "Unknown instruction")
        }
    }
}

func intCall(ctx *Context, ins node.Node) {
    foo := ins.Children[1].(token.Token)

    switch foo.Str {
    case "printc": fmt.Printf("%c", ctx.GetRegByNum(0))
    case "printi": fmt.Print(ctx.GetRegByNum(0))
    default:
        util.Fail(foo, "Unknown function called")
    }
}

func intMove(ctx *Context, ins node.Node) {
    val := ins.Children[1].(node.Node)
    reg := ins.Children[3].(node.Node)

    v := intValue(ctx, val)
    ctx.SetReg(reg, v)
}

func intAdd(ctx *Context, ins node.Node) {
    val := ins.Children[1].(node.Node)
    reg := ins.Children[3].(node.Node)

    v := intValue(ctx, val)
    v += ctx.GetReg(reg)
    ctx.SetReg(reg, v)
}

func intSub(ctx *Context, ins node.Node) {
    val := ins.Children[1].(node.Node)
    reg := ins.Children[3].(node.Node)

    v := -intValue(ctx, val)
    v += ctx.GetReg(reg)
    ctx.SetReg(reg, v)
}

func intInc(ctx *Context, ins node.Node) {
    reg := ins.Children[1].(node.Node)
    v := ctx.GetReg(reg)
    ctx.SetReg(reg, v + 1)
}

func intDec(ctx *Context, ins node.Node) {
    reg := ins.Children[1].(node.Node)
    v := ctx.GetReg(reg)
    ctx.SetReg(reg, v - 1)
}

func intValue(ctx *Context, val node.Node) uint64 {
    if val.Name == "Reg" { return ctx.GetReg(val) }

    s := val.Children[0].(token.Token).Str
    v, err := strconv.ParseUint(s, 10, 64)
    if err != nil { util.Fail(val, "Bad value") }

    return v
}

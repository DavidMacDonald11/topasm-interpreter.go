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
        case "label": continue
        case "jump": intJump(&ctx, ins)
        case "call": intCall(&ctx, ins)
        case "move": intMove(&ctx, ins)
        case "add": intAdd(&ctx, ins)
        case "sub": intSub(&ctx, ins)
        case "comp": intComp(&ctx, ins)
        case "inc": intInc(&ctx, ins)
        case "dec": intDec(&ctx, ins)
        case "mul": intMul(&ctx, ins)
        case "div": intDiv(&ctx, ins)
        default: util.Fail(ins, "Unknown instruction")
        }
    }
}

func intJump(ctx *Context, ins node.Node) {
    jump := ins.Children[0].(token.Token)
    label := ins.Children[1].(token.Token)

    l := ctx.GetLabel(label)
    eq, lt := ctx.GetCompFlags()

    switch jump.Str {
    case "jump":
    case "jumpNE": if !eq { ctx.Jump(l) }
    case "jumpEQ": if eq { ctx.Jump(l) }
    case "jumpLT": if lt { ctx.Jump(l) }
    case "jumpGT": if !lt && !eq { ctx.Jump(l) }
    case "jumpLTE": if lt || eq { ctx.Jump(l) }
    case "jumpGTE": if !lt { ctx.Jump(l) }
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

func intComp(ctx *Context, ins node.Node) {
    val1 := ins.Children[1].(node.Node)
    val2 := ins.Children[3].(node.Node)

    v1 := intValue(ctx, val1)
    v2 := intValue(ctx, val2)

    ctx.Comp(v1, v2)
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

func intMul(ctx *Context, ins node.Node) {
    reg := ins.Children[1].(node.Node)
    n := ctx.GetRegByNum(0)
    q := ctx.GetReg(reg)

    ctx.SetRegByNum(0, n * q)
}

func intDiv(ctx *Context, ins node.Node) {
    reg := ins.Children[1].(node.Node)
    n := ctx.GetRegByNum(0)
    q := ctx.GetReg(reg)

    ctx.SetRegByNum(0, n / q)
    ctx.SetRegByNum(1, n % q)
}

func intValue(ctx *Context, val node.Node) uint64 {
    if val.Name == "reg" { return ctx.GetReg(val) }

    s := val.Children[0].(token.Token).Str
    v, err := strconv.ParseUint(s, 10, 64)
    if err != nil { util.Fail(val, "Bad value") }

    return v
}

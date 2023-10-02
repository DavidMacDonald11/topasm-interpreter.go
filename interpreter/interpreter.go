package interpreter

import (
	"log"
	"strconv"
	"topasm/fault"
	"topasm/node"
	"topasm/token"
)

func InterpretTree(tree *node.Node) *fault.Fault {
    ctx := NewContext()
    var f *fault.Fault

    for _, child := range tree.Children {
        if ins, ok := child.(*node.Node); ok {
            switch ins.Kind {
            case "Binary Instruction": f = interpretBinary(ctx, ins)
            case "Unary Instruction": f = interpretUnary(ctx, ins)
            default: log.Fatal("Bad node kind")
            }
        } else { continue }

        if f != nil { return f }
    }

    return nil
}

func interpretBinary(ctx *Context, ins *node.Node) *fault.Fault {
    return nil
}

func interpretUnary(ctx *Context, ins *node.Node) *fault.Fault {
    name, _ := ins.Children[0].(*token.Token)
    valNode, _ := ins.Children[1].(*node.Node)

    val, f := getValue(ctx, valNode)
    if f != nil { return f }

    if name.Str == "printc" { ctx.Printc(*val) } else { ctx.Printi(*val) }
    return nil
}

func getValue(ctx *Context, val *node.Node) (*uint64, *fault.Fault) {
    if val.Kind == "Num" {
        tok, _ := val.Children[0].(*token.Token)
        num, err := strconv.ParseUint(tok.Str, 10, 64)
        if err != nil { log.Fatal(err) }

        return &num, nil
    }

    tok, _ := val.Children[1].(*token.Token)
    reg, fault := ctx.VerifyReg(tok)

    if fault != nil { return nil, fault }

    num := ctx.ReadReg(reg)
    return &num, nil
}

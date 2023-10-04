package interpreter

import (
	"fmt"
	"log"
	"strconv"
	"topasm/fault"
	"topasm/node"
)

func InterpretTree(tree node.File) {
    ctx := NewContext()

    for i := 0; i < len(tree.Instructions); i += 2 {
        ins := tree.Instructions[i]

        unary, ok := ins.(node.UnaryIns)
        if ok { interpretUnaryIns(&ctx, unary); continue }

        binary, ok := ins.(node.BinaryIns)
        if ok { interpretBinaryIns(&ctx, binary); continue }

        fault.Fail(ins, "Interpreting", "Unexpected instruction")
    }
}

func interpretUnaryIns(ctx *Context, ins node.UnaryIns) {
    val := ctx.GetReg(ins.Reg)

    if ins.Ins.Has("inc") {
        ctx.SetReg(ins.Reg, val + 1)
    } else if ins.Ins.Has("dec") {
        ctx.SetReg(ins.Reg, val - 1)
    } else if ins.Ins.Has("printc") {
        fmt.Printf("%c", val)
    } else {
        print(val)
    }
}

func interpretBinaryIns(ctx *Context, ins node.BinaryIns) {
    var val uint64

    if r, ok := ins.Val.(node.Reg); ok {
        val = ctx.GetReg(r)
    } else if n, ok := ins.Val.(node.Num); ok {
        var err error
        val, err = strconv.ParseUint(n.Num.Str, 10, 64)
        if err != nil { log.Fatal(err) }
    } else {
        fault.Fail(ins.Val, "Interpreting", "Unexpected value")
    }

    if ins.Ins.Has("move") {
        ctx.SetReg(ins.Reg, val)
    } else if ins.Ins.Has("add") {
        val += ctx.GetReg(ins.Reg)
        ctx.SetReg(ins.Reg, val)
    } else if ins.Ins.Has("sub") {
        val = -val
        val += ctx.GetReg(ins.Reg)
        ctx.SetReg(ins.Reg, val)
    }
}

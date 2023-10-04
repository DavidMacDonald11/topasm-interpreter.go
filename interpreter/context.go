package interpreter

import (
	"log"
	"strconv"
	"topasm/fault"
	"topasm/node"
)

type Context struct {
    regs [10]uint64
}

func NewContext() Context {
    return Context{[10]uint64{}}
}

func (c *Context) SetReg(r node.Reg, val uint64) {
    reg := verifyReg(r)
    c.regs[reg] = val
}

func (c *Context) GetReg(r node.Reg) uint64 {
    reg := verifyReg(r)
    return c.regs[reg]
}

func verifyReg(r node.Reg) int {
    reg, err := strconv.Atoi(r.Num.Str)
    if err != nil { log.Fatal(err) }
    if reg > 9 { fault.Fail(r, "Interpreting", "No such register") }

    return reg
}

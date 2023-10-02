package interpreter

import (
	"strconv"
	"topasm/fault"
	"topasm/token"
)

type Context struct {
    reg [10]uint64
}

func NewContext() *Context {
    return &Context{[10]uint64{}}
}

func (c *Context) VerifyReg(reg *token.Token) (int, *fault.Fault) {
    num, err := strconv.Atoi(reg.Str)
    if err == nil && num < 10 { return num, nil }
    return num, fault.New(reg, "Interpreting", "There is no such register")
}

func (c *Context) ReadReg(reg int) uint64 {
    return c.reg[reg]
}

func (c *Context) Move(val uint64, reg int) {
    c.reg[reg] = val
}

func (c *Context) Add(val uint64, reg int) {
    c.reg[reg] += val
}

func (c *Context) Sub(val uint64, reg int) {
    c.reg[reg] -= val
}

func (c *Context) Printc(val uint64) {
    print(rune(val))
}

func (c *Context) Printi(val uint64) {
    print(val)
}

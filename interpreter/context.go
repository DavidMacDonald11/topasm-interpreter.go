package interpreter

import (
	"strconv"
	"topasm/node"
	"topasm/token"
	"topasm/util"
)

type Context struct {
    i int
    regs [10]uint64
    labels map[string]int
    eqFlag bool
    ltFlag bool
}

func NewContext(tree node.Node) Context {
    labels := make(map[string]int)

    for i := 0; i * 2 < len(tree.Children); i++ {
        ins := tree.Children[i * 2].(node.Node)
        if ins.Name != "label" { continue }

        labels[ins.Children[0].(token.Token).Str] = i
    }

    return Context{0, [10]uint64{}, labels, false, false}
}

func (c *Context) SetReg(r node.Node, val uint64) {
    reg := c.verifyReg(r)
    c.regs[reg] = val
}

func (c *Context) GetReg(r node.Node) uint64 {
    reg := c.verifyReg(r)
    return c.regs[reg]
}

func (c *Context) GetRegByNum(r int) uint64 {
    return c.regs[r]
}

func (c *Context) Comp(n1 uint64, n2 uint64) {
    c.eqFlag = (n1 == n2)
    c.ltFlag = (n1 < n2)
}

func (c *Context) GetCompFlags() (bool, bool) {
    return c.eqFlag, c.ltFlag
}

func (c *Context) GetLabel(label token.Token) string {
    _, ok := c.labels[label.Str]
    if !ok { util.Fail(label, "No such label") }

    return label.Str
}

func (c *Context) Jump(label string) {
    c.i = c.labels[label]
}

func (c *Context) verifyReg(r node.Node) int {
    reg, err := strconv.Atoi(r.Children[1].(token.Token).Str)
    if err != nil || reg + 1 >= len(c.regs) { util.Fail(r, "No such register") }

    return reg
}

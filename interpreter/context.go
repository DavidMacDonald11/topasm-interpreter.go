package interpreter

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"topasm/node"
	"topasm/token"
	"topasm/util"
)

type counter struct {
    label string
    n uint64
}

type Context struct {
    i int
    regs [10]uint64
    labels map[string]int
    eqFlag bool
    ltFlag bool
    loopCounter counter
}

func NewContext(tree node.Node) Context {
    labels := make(map[string]int)

    for i := 0; i * 2 < len(tree.Children); i++ {
        ins := tree.Children[i * 2].(node.Node)
        if ins.Name != "label" { continue }

        tok := ins.Children[0].(token.Token)

        if _, ok := labels[tok.Str]; ok {
            util.Fail(tok, "Duplicate label")
        }

        labels[tok.Str] = i
    }

    return Context{0, [10]uint64{}, labels, false, false, counter{}}
}

func (c *Context) SetReg(r node.Node, val uint64) {
    reg := c.verifyReg(r)
    c.regs[reg] = val
}

func (c *Context) SetRegByNum(r int, val uint64) {
    c.regs[r] = val
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

    if c.loopCounter.label != label {
        c.loopCounter = counter{label, 0}
    } else if c.loopCounter.n > 499999999 {
        fmt.Printf("Infinte loop detected at label '%s'\n", label)
        os.Exit(2)
    } else {
        c.loopCounter.n++
    }
}

func (c Context) String() string {
    b := strings.Builder{}
    b.WriteRune('[')

    for i, r := range c.regs {
        b.WriteString(fmt.Sprintf("%d", r))
        if i != len(c.regs) - 1 { b.WriteString(", ") }
    }

    b.WriteString("], eq: ")
    b.WriteRune(util.IfElse(c.eqFlag, '1', '0'))
    b.WriteString(", lt: ")
    b.WriteRune(util.IfElse(c.ltFlag, '1', '0'))

    return b.String()
}

func (c *Context) verifyReg(r node.Node) int {
    reg, err := strconv.Atoi(r.Children[1].(token.Token).Str)
    if err != nil || reg >= len(c.regs) { util.Fail(r, "No such register") }

    return reg
}

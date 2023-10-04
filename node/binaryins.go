package node

import (
	"topasm/fault"
	"topasm/token"
)

type Token = token.Token

type BinaryIns struct {
    Ins token.Token
    Val Node
    prep token.Token
    Reg Reg
}

func NewBinaryIns(ins Token, val Node, prep Token, reg Reg) BinaryIns {
    return BinaryIns{ins, val, prep, reg}
}

func (b BinaryIns) Position() fault.Position {
    start := b.Ins.Position().Start
    end := b.Reg.Position().End
    return fault.NewPosition(start, end)
}

func (b BinaryIns) NodeString(prefix string) string {
    return nodeString(prefix, "BinaryIns", b.Ins, b.Val, b.prep, b.Reg)
}

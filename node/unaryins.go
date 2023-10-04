package node

import (
	"topasm/fault"
	"topasm/token"
)

type UnaryIns struct {
    Ins token.Token
    Reg Reg
}

func NewUnaryIns(ins token.Token, reg Reg) UnaryIns {
    return UnaryIns{ins, reg}
}

func (u UnaryIns) Position() fault.Position {
    start := u.Ins.Position().Start
    end := u.Reg.Position().End
    return fault.NewPosition(start, end)
}

func (u UnaryIns) NodeString(prefix string) string {
    return nodeString(prefix, "UnaryIns", u.Ins, u.Reg)
}

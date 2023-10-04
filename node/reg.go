package node

import (
	"topasm/token"
	"topasm/fault"
)

type Reg struct {
    hash token.Token
    Num token.Token
}

func NewReg(hash token.Token, num token.Token) Reg {
    return Reg{hash, num}
}

func (r Reg) Position() fault.Position {
    start := r.hash.Position().Start
    end := r.Num.Position().End
    return fault.NewPosition(start, end)
}

func (r Reg) NodeString(prefix string) string {
    return nodeString(prefix, "Reg", r.hash, r.Num)
}

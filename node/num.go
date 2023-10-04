package node

import (
	"topasm/token"
	"topasm/fault"
)

type Num struct {
    Num token.Token
}

func NewNum(tok token.Token) Num {
    return Num{tok}
}

func (n Num) Position() fault.Position {
    return n.Num.Position()
}

func (n Num) NodeString(prefix string) string {
    return nodeString(prefix, "Num", n.Num)
}

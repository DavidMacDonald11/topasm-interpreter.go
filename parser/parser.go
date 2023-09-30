package parser

import (
	"topasm/core"
	"topasm/core/node"
	"topasm/core/token"
)

type Tokens = token.Tokens
type Fault = core.Fault
type Faults = core.Faults
type Node = node.Node

func ParseTokens(tokens Tokens) Node {
    return node.NewNode("Test", &tokens[0])
}

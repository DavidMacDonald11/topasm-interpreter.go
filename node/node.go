package node

import (
	"fmt"
	"strings"
	"topasm/fault"
	"topasm/util"
)

type Node interface {
    NodeString(string) string
    Position() fault.Position
}

func nodeString(prefix string, name string, children ...Node) string {
    b := strings.Builder{}
    b.WriteString(name)

    for i, child := range children {
        isLast := i == len(children) - 1

        branch := util.IfElse(isLast, " ", "│")
        childPrefix := fmt.Sprintf("%s%s   ", prefix, branch)
        childStr := child.NodeString(childPrefix)

        branch = util.IfElse(isLast, "└──", "├──")
        str := fmt.Sprintf("\n%s%s %s", prefix, branch, childStr)
        b.WriteString(str)
    }

    return b.String()
}

type Error struct {
    node Node
}

func NewError(node Node) Error {
    return Error{node}
}

func (e Error) Position() fault.Position {
    return e.Position()
}

func (e Error) NodeString(prefix string) string {
    return nodeString(prefix, "?", e.node)
}

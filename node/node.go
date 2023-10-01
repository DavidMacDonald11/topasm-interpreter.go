package node

import (
	"fmt"
	"strings"
	"topasm/fault"
	"topasm/util"
)

type Child interface {
    Position() fault.Position
    String() string
}

type Node struct {
    Kind string
    Children []Child
}

func New(kind string, children ...Child) *Node {
    return &Node{Kind: kind, Children: children}
}

func (n Node) Position() fault.Position {
    first := n.Children[0].Position()
    last := n.Children[len(n.Children) - 1].Position()
    return *fault.NewPosition(first.Start, last.End)
}

func (n Node) String() string {
    return n.treeString("")
}

func (n Node) treeString(prefix string) string {
    b := strings.Builder{}
    b.WriteString(n.Kind)

    for i, child := range n.Children {
        isLast := i == len(n.Children) - 1
        var childStr string

        if node, ok := child.(*Node); ok {
            branch := util.IfElse(isLast, " ", "│")
            pre := fmt.Sprintf("%s%s   ", prefix, branch)
            childStr = node.treeString(pre)
        } else { childStr = child.String() }

        branch := util.IfElse(isLast, "└──", "├──")
        str := fmt.Sprintf("\n%s%s %s", prefix, branch, childStr)
        b.WriteString(str)
    }

    return b.String()
}

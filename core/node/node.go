package node

import (
	"fmt"
	"strings"
	"topasm/core"
)

type child interface {
    Position() core.UIntRange
    String() string
}

type Node struct {
    kind string
    children []child
}

func NewNode(kind string, children ...child) Node {
    return Node {
        kind: kind,
        children: children,
    }
}

func (n Node) Position() core.UIntRange {
    return core.UIntRange {
        Start: n.children[0].Position().Start,
        End: n.children[len(n.children) - 1].Position().End,
    }
}

func (n Node) String() string {
    return n.treeString("")
}

func (n Node) treeString(prefix string) string {
    builder := strings.Builder{}
    builder.WriteString(n.kind)

    for i, child := range n.children {
        isLast := i == len(n.children) - 1
        var childStr string

        if _, ok := child.(Node); ok {
            branch := core.IfElse(isLast, " ", "│")
            pre := fmt.Sprintf("%s%s   ", prefix, branch)
            childStr = child.(Node).treeString(pre)
        } else { childStr = child.String() }

        branch := core.IfElse(isLast, "└──", "├──")
        str := fmt.Sprintf("\n%s%s %s", prefix, branch, childStr)
        builder.WriteString(str)
    }

    return builder.String()
}

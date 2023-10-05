package node

import (
	"fmt"
	"strings"
	"topasm/util"
)

type RecStringer interface {
    Position() int
    RecString(string) string
}

type Node struct {
    Name string
    Children []RecStringer
}

func New(name string, children ...RecStringer) Node {
    return Node{name, children}
}

func (n Node) Position() int {
    return n.Children[0].Position()
}

func (n Node) String() string {
    return n.RecString("")
}

func (n Node) RecString(prefix string) string {
    b := strings.Builder{}
    b.WriteString(n.Name)

    for i, pair := range n.Children {
        isLast := i == len(n.Children) - 1

        branch := util.IfElse(isLast, " ", "│")
        childPrefix := fmt.Sprintf("%s%s   ", prefix, branch)
        childStr := pair.RecString(childPrefix)

        branch = util.IfElse(isLast, "└──", "├──")
        str := fmt.Sprintf("\n%s%s %s", prefix, branch, childStr)
        b.WriteString(str)
    }

    return b.String()
}

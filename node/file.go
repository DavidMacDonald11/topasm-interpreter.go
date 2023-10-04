package node

import "topasm/fault"

type File struct {
    Instructions []Node
}

func NewFile(instructions []Node) File {
    return File{instructions}
}

func (f File) Position() fault.Position {
    start := f.Instructions[0].Position().Start
    end := f.Instructions[len(f.Instructions)].Position().End
    return fault.NewPosition(start, end)
}

func (f File) NodeString(prefix string) string {
    return nodeString(prefix, "File", f.Instructions...)
}

func (f File) String() string {
    return f.NodeString("")
}

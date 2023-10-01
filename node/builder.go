package node

import (
	"topasm/fault"
)

type Builder struct {
    kind string
    children []Child
    fault *fault.Fault
}

func NewBuilder(kind string) *Builder {
    return &Builder{kind, []Child{}, nil}
}

func (b *Builder) HasFault() bool { return b.fault != nil }

func (b *Builder) AddChild(c Child) {
    if c != nil { b.children = append(b.children, c) }
}

func (b *Builder) AddFault(f *fault.Fault) {
    if b.fault == nil { b.fault = f }
}

func (b *Builder) Add(c Child, f *fault.Fault) {
    b.AddChild(c)
    b.AddFault(f)
}

func (b *Builder) AddResult(r *Result) {
    b.AddChild(r.Node)
    b.AddFault(r.Fault)
}

func (b *Builder) Result() *Result {
    node := New(b.kind, b.children...)
    return &Result{node, b.fault}
}

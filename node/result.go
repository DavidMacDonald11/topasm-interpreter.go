package node

import "topasm/fault"

type Result struct {
    Node *Node
    Fault *fault.Fault
}

func (r *Result) HasNode() bool { return r.Node != nil }
func (r *Result) HasFault() bool { return r.Fault != nil }

func NoneResult() *Result { return &Result{nil, nil} }
func NodeResult(n *Node) *Result { return &Result{n, nil} }

func FaultResult(n *Node, label string, msg string) *Result {
    return &Result{n, fault.New(n, label, msg)}
}

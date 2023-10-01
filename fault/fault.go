package fault

import "fmt"

type Position struct {
    Start int
    End int
}

func NewPosition(start int, end int) *Position {
    return &Position{Start: start, End: end}
}

type Positioner interface {
    Position() Position
}

type Fault struct {
    Pos Position
    Msg string
    IsFail bool
}

func NewFault(p Positioner, label string, msg string, isFail bool) *Fault {
    msg = fmt.Sprintf("%s Error: %s", label, msg)
    return &Fault{p.Position(), msg, isFail}
}

func (f *Fault) Print() {
    fmt.Printf("%s (at %d:%d)\n", f.Msg, f.Pos.Start, f.Pos.End)
}

package fault

import (
	"log"
)

type Position struct {
    Start, End int
}

func NewPosition(start, end int) Position {
    return Position{start, end}
}

type Positioner interface {
    Position() Position
}

func Fail(p Positioner, label string, msg string) {
    pos := p.Position()
    log.Fatalf("%s Error: %s (at %d:%d)\n", label, msg, pos.Start, pos.End)
}

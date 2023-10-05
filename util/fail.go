package util

import "log"

type Positioner interface {
    Position() int
}

func Fail(p Positioner, msg string) {
    pos := p.Position()
    log.Fatalf("Error: %s (on line %d)\n", msg, pos)
}

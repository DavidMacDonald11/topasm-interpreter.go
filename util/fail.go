package util

import (
	"fmt"
	"os"
)

type Positioner interface {
    Position() int
}

func Fail(p Positioner, msg string) {
    pos := p.Position()
    fmt.Printf("Error: %s (on line %d)\n", msg, pos)
    os.Exit(1)
}

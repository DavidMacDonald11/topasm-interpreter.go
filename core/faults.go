package core

import (
	"fmt"
	"log"
	"strings"
)

type position = UIntRange

type filePos struct {
    char uint64
    line int64
}

type Positioner interface {
    Position() position
}

type Fault struct {
    pos position
    msg string
    isFail bool
}

func NewFault(p Positioner, label string, msg string, isFail bool) Fault {
    return Fault {
        pos: p.Position(),
        msg: fmt.Sprintf("%s Error: %s", label, msg),
        isFail: isFail,
    }
}

func (f Fault) IsFail() bool { return f.isFail }

func (f Fault) print(file FileLines) {
    println(f.msg)

    pos := filePos{}
    f.seek(file, &pos)

    for {
        if pos.line == int64(len(file)) {
            fmt.Printf("%4d|EOF\n    |^^^\n", pos.line + 1)
            return
        }

        line := file[pos.line]
        marks := strings.Builder{}

        for range line {
            c := IfElse(f.pos.contains(pos.char), '^', ' ')
            marks.WriteRune(c)
            pos.char += uint64(1)
        }

        fmt.Printf("%4d|%s    |%s\n", pos.line + 1, line, marks.String())
        pos.line += 1

        if pos.char >= f.pos.End { break }
    }
}

func (f Fault) seek(file FileLines, pos *filePos) {
    for _, line := range file {
        char := pos.char + uint64(len(line))
        if char > f.pos.Start { break }

        pos.char = char
        pos.line += 1
    }
}

type Faults []Fault

func (f Faults) Print(filePath string) {
    file, err := NewFileLines(filePath)
    if err != nil { log.Fatal(err) }

    for _, fault := range f {
        fault.print(file)
    }
}

package core

import (
	"fmt"
	"os"
	"strings"
)

type UIntRange struct {
    Start uint64
    End uint64
}

func (self *UIntRange) contains(n uint64) bool {
    return n >= self.Start && n <= self.End
}

type Faultable interface {
    FaultPosition() UIntRange
}

type Fault struct {
    Obj Faultable
    Message string
}

func NewFault(obj Faultable, label string, message string) Fault {
    return Fault {
        Obj: obj,
        Message: fmt.Sprintf("%s Error: %s", label, message),
    }
}

func PrintFaults(fileName string, faults []Fault) error {
    fileLines, err := readFile(fileName)
    if err != nil { return err }

    for _, fault := range faults {
        printFault(fileLines, fault)
    }

    return nil
}

func readFile(fileName string) ([]string, error) {
    file, err := os.ReadFile(fileName)
    if err != nil { return nil, err }

    var fileLines []string

    for _, line := range strings.Split(string(file), "\n") {
        fileLines = append(fileLines, line + "\n")
    }

    return fileLines, nil
}

type fileCounter struct {
    chars uint64
    lines int64
}

func printFault(fileLines []string, fault Fault) {
    position := fault.Obj.FaultPosition()
    counter := fileCounter {}

    seek(fileLines, fault, &counter)

    println(fault.Message)
    printLine(fileLines, fault, &counter)

    for counter.chars < position.End {
        printLine(fileLines, fault, &counter)
    }
}

func seek(fileLines []string, fault Fault, counter *fileCounter) {
    position := fault.Obj.FaultPosition()

    for _, line := range fileLines {
        lineLen := uint64(len(line))
        if counter.chars + lineLen > position.Start { break }

        counter.chars += lineLen
        counter.lines++
    }
}

func printLine(fileLines []string, fault Fault, counter *fileCounter) {
    if counter.lines == int64(len(fileLines)) {
        fmt.Printf("%4d|EOF\n    |^^^\n", counter.lines + 1)
        return
    }

    position := fault.Obj.FaultPosition()
    line := fileLines[counter.lines]
    marks := strings.Builder {}

    for range line {
        c := IfThen(position.contains(counter.chars), '^', ' ')
        marks.WriteRune(c)
        counter.chars++
    }

    fmt.Printf("%4d|%s    |%s\n", counter.lines + 1, line, marks.String())
    counter.lines++
}

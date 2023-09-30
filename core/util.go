package core

import (
	"os"
	"strings"
)

func IfElse[T any](c bool, val1 T, val2 T) T {
    if c { return val1 }
    return val2
}

func Map[T any](slice []T, f func(T) T) []T {
    var result []T

    for _, t := range slice {
        result = append(result, f(t))
    }

    return result
}

type UIntRange struct {
    Start uint64
    End uint64
}

func (u *UIntRange) contains(n uint64) bool {
    return n >= u.Start && n <= u.End
}

type FileLines []string

func NewFileLines(path string) (FileLines, error) {
    file, err := os.ReadFile(path)
    if err != nil { return nil, err }

    lines := strings.Split(string(file), "\n")
    lines = Map(lines, func(s string) string { return s + "\n" })
    return FileLines(lines), nil
}

package io

import (
	"log"
	"os"
	"strings"
    "topasm/core"
)

type SourceFile struct {
    path string
    contents string
    charPos uint64
    buffer strings.Builder
}

func NewSourceFile(path string) SourceFile {
    file, err := os.ReadFile(path)
    if err != nil { log.Fatal(err) }

    return SourceFile {
        path: path,
        contents: string(file),
        charPos: uint64(0),
        buffer: strings.Builder{},
    }
}

func (s *SourceFile) CharPos() uint64 {
    return s.charPos
}

func (s *SourceFile) AtEnd() bool {
    return s.charPos == s.chars() - uint64(1)
}

func (s *SourceFile) Peek(n int) string {
    lower := s.charPos + uint64(n - 1)

    if lower >= s.chars() { return "" }
    return s.contents[lower:]
}

func (s *SourceFile) NextChar() rune {
    return core.IfElse(s.AtEnd(), '\u0000', rune(s.Peek(1)[0]))
}

func (s *SourceFile) ReadChar() string {
    return s.ReadChars(1)
}

func (s *SourceFile) ReadChars(n int) string {
    return s.readWhile(func() bool {
        return len(s.buffer.String()) < n
    })
}

func (s *SourceFile) ReadTheseChars(these string) string {
    return s.readWhile(func() bool {
        return strings.ContainsRune(these, s.NextChar())
    })
}

func (s *SourceFile) ReadCharsUntil(until string) string {
    return s.readWhile(func() bool {
        return !strings.ContainsRune(until, s.NextChar())
    })
}

func (s *SourceFile) readWhile(predicate func() bool) string {
    s.buffer.Reset()

    for !s.AtEnd() && predicate() {
        s.buffer.WriteRune(s.NextChar())
        s.charPos += uint64(1)
    }

    return s.buffer.String()
}

func (s *SourceFile) chars() uint64 { return uint64(len(s.contents)) }

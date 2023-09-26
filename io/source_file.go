package io

import (
	"log"
	"os"
	"strings"
    "topasm/core"
)

type SourceFile struct {
    Path string
    contents string
    charPos uint64
    buffer strings.Builder
}

func NewSourceFile(path string) SourceFile {
    file, err := os.ReadFile(path)
    if err != nil { log.Fatal(err) }

    return SourceFile {
        Path: path,
        contents: string(file),
        charPos: uint64(0),
        buffer: strings.Builder {},
    }
}

func (self *SourceFile) CharPos() uint64 {
    return self.charPos
}

func (self *SourceFile) AtEnd() bool {
    return self.charPos == self.chars() - uint64(1)
}

func (self *SourceFile) Peek(n int) string {
    lower := self.charPos + uint64(n - 1)

    if lower >= self.chars() { return "" }
    return self.contents[lower:]
}

func (self *SourceFile) NextChar() byte {
    return core.IfThen(self.AtEnd(), '\u0000', self.Peek(1)[0])
}

func (self *SourceFile) ReadChar() string {
    return self.ReadChars(1)
}

func (self *SourceFile) ReadChars(n int) string {
    return self.readWhile(func() bool {
        return len(self.buffer.String()) < n
    })
}

func (self *SourceFile) ReadTheseChars(these string) string {
    return self.readWhile(func() bool {
        return strings.ContainsRune(these, rune(self.NextChar()))
    })
}

func (self *SourceFile) ReadCharsUntil(until string) string {
    return self.readWhile(func() bool {
        return !strings.ContainsRune(until, rune(self.NextChar()))
    })
}

func (self *SourceFile) readWhile(predicate func() bool) string {
    self.buffer.Reset()

    for !self.AtEnd() && predicate() {
        self.buffer.WriteByte(self.NextChar())
        self.charPos++
    }

    return self.buffer.String()
}

func (self *SourceFile) chars() uint64 { return uint64(len(self.contents)) }

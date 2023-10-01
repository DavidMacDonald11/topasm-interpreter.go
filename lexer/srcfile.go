package lexer

import (
	"log"
	"os"
	"strings"
	"topasm/util"
)

type SrcFile struct {
    contents string
    buf strings.Builder
    Pos int
}

func NewSrcFile(path string) *SrcFile {
    file, err := os.ReadFile(path)
    if err != nil { log.Fatal(err) }

    return &SrcFile {
        contents: string(file),
        buf: strings.Builder{},
        Pos: 0,
    }
}

func (s *SrcFile) AtEnd() bool { return s.Pos == len(s.contents) - 1 }

func (s *SrcFile) NextChar() rune {
    return util.IfElse(s.AtEnd(), '\u0000', rune(s.Peek(1)[0]))
}

func (s *SrcFile) Peek(n int) string {
    n += s.Pos - 1
    return util.IfElse(n >= len(s.contents), "", s.contents[n:])
}

func (s *SrcFile) ReadChar() string { return s.ReadChars(1) }

func (s *SrcFile) ReadChars(n int) string {
    return s.readWhile(func() bool {
        return len(s.buf.String()) < n
    })
}

func (s *SrcFile) ReadTheseChars(these string) string {
    return s.readWhile(func() bool {
        return strings.ContainsRune(these, s.NextChar())
    })
}

func (s *SrcFile) ReadCharsUntil(until string) string {
    return s.readWhile(func() bool {
        return !strings.ContainsRune(until, s.NextChar())
    })
}

func (s *SrcFile) readWhile(predicate func() bool) string {
    s.buf.Reset()

    for !s.AtEnd() && predicate() {
        s.buf.WriteRune(s.NextChar())
        s.Pos += 1
    }

    return s.buf.String()
}

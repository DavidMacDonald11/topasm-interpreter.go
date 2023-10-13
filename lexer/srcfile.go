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
    pos int
    Line int
}

func OpenSrcFile(path string) SrcFile {
    file, err := os.ReadFile(path)
    if err != nil { log.Fatal(err) }

    return SrcFile {
        contents: string(file) + "\n",
        buf: strings.Builder{},
        pos: 0,
        Line: 1,
    }
}

func MakeSrcFile(text string) SrcFile {
    return SrcFile {
        contents: text + "\n",
        buf: strings.Builder{},
        pos: 0,
        Line: 1,
    }
}

func (s SrcFile) AtEnd() bool { return s.pos == len(s.contents) - 1 }

func (s SrcFile) NextChar() rune {
    return util.IfElse(s.AtEnd(), '\u0000', rune(s.Peek(1)[0]))
}

func (s SrcFile) Peek(n int) string {
    n += s.pos - 1
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
        c := s.NextChar()
        if c == '\n' { s.Line += 1 }

        s.buf.WriteRune(c)
        s.pos += 1
    }

    return s.buf.String()
}

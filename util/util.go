package util

import (
	"fmt"
	"strings"
)

func IfElse[T any](c bool, val1 T, val2 T) T {
    if c { return val1 } else { return val2 }
}

func Join[T fmt.Stringer](s []T, sep string, pre string, post string) string {
    builder := strings.Builder{}
    builder.WriteString(pre)

    for i, t := range s {
        if i != 0 { builder.WriteString(sep) }
        builder.WriteString(t.String())
    }

    builder.WriteString(post)
    return builder.String()
}

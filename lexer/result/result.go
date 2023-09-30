package result

import (
    "topasm/core"
    "topasm/core/token"
)

type Result struct {
    Token *token.Token
    Fault *core.Fault
}

func (r *Result) HasToken() bool {
    return r.Token != nil
}

func (r *Result) HasFault() bool {
    return r.Fault != nil
}

func (r *Result) Failed() bool {
    return r.Fault != nil && r.Fault.IsFail()
}

func None() Result {
    return Result{nil, nil}
}

func Token(t token.Token) Result {
    return Result{&t, nil}
}

func Error(t token.Token, msg string) Result {
    fault := core.NewFault(&t, "Lexing", msg, false)
    return Result{&t, &fault}
}

func Failure(t token.Token, msg string) Result {
    fault := core.NewFault(&t, "Lexing", msg, true)
    return Result{&t, &fault}
}

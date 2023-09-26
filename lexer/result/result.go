package result

import (
    "topasm/core"
    "topasm/core/token"
)

type Result struct {
    Token *token.Token
    Fault *core.Fault
}

func (self *Result) HasToken() bool {
    return self.Token != nil
}

func (self *Result) HasFault() bool {
    return self.Fault != nil
}

func (self *Result) Failed() bool {
    return self.Fault != nil && self.Fault.Fail
}

func None() Result {
    return Result{nil, nil}
}

func Token(token token.Token) Result {
    return Result{&token, nil}
}

func Error(token token.Token, message string) Result {
    fault := core.NewFault(&token, "Lexing", message, false)
    return Result{&token, &fault}
}

func Failure(token token.Token, message string) Result {
    fault := core.NewFault(&token, "Lexing", message, true)
    return Result{&token, &fault}
}

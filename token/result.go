package token

import "topasm/fault"

type Result struct {
    Token *Token
    Fault *fault.Fault
}

func (r *Result) HasToken() bool { return r.Token != nil }
func (r *Result) HasFault() bool { return r.Fault != nil }
func (r *Result) Failed() bool { return r.HasFault() && r.Fault.IsFail }

func NoneResult() *Result { return &Result{nil, nil} }
func TokenResult(t *Token) *Result { return &Result{t, nil} }

func ErrorResult(t *Token, label string, msg string) *Result {
    return &Result {
        Token: t,
        Fault: fault.NewFault(t, label, msg, false),
    }
}

func FailureResult(t *Token, label string, msg string) *Result {
    return &Result {
        Token: t,
        Fault: fault.NewFault(t, label, msg, true),
    }
}

package token

import "topasm/fault"

type Result struct {
    Token *Token
    Fault *fault.Fault
}

func (r *Result) HasToken() bool { return r.Token != nil }
func (r *Result) HasFault() bool { return r.Fault != nil }

func NoneResult() *Result { return &Result{nil, nil} }
func TokenResult(t *Token) *Result { return &Result{t, nil} }

func FaultResult(t *Token, label string, msg string) *Result {
    return &Result{t, fault.New(t, label, msg)}
}

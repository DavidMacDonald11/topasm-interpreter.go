package grammar

const (
    EOF = ""
    Digits = "0123456789"
    Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
    Puncs = "\n:#"
)

func Keys() []string {
    return append([]string {
        "move", "into", "add", "sub", "from", "comp", "with",
        "inc", "dec", "call",
    }, JumpKeys()...)
}

func JumpKeys() []string {
    return []string {
        "jump", "jumpNE", "jumpEQ", "jumpLT", "jumpGT", "jumpLTE", "jumpGTE",
    }
}

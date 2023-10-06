package grammar

const (
    EOF = ""
    Digits = "0123456789"
    Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
    Puncs = "\n:#"
    EscapeSymbols = "'\"\\0abfnrtv"
)

func Keys() []string {
    return append([]string {
        "move", "into", "add", "sub", "from", "comp", "with",
        "mul", "div",
        "inc", "dec", "call",
    }, JumpKeys()...)
}

func JumpKeys() []string {
    return []string {
        "jump", "jumpNE", "jumpEQ", "jumpLT", "jumpGT", "jumpLTE", "jumpGTE",
    }
}

func EscapeSymbolMap() map[byte]uint64 {
    return map[byte]uint64{
        '\'': 39, '"': 34, '\\': 92, '0': 0, 'a': 7,
        'b': 8, 'f': 12, 'n': 10, 'r': 13, 't': 9, 'v': 11,
    }
}

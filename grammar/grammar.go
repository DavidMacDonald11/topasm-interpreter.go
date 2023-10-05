package grammar

const (
    EOF = ""
    Digits = "0123456789"
    Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
    Puncs = "\n:#"
)

func Keys() []string {
    return []string {
        "move", "into", "add", "sub", "from",
        "inc", "dec", "call",
    }
}

package grammar

const (
    EOF = ""
    Digits = "0123456789"
    Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
    Puncs = ":#"
)

func Keys() []string {
    return []string {
        "move", "into", "add", "sub", "from",
        "printc", "printi",
    }
}

package grammar

const (
    EOF = ""
    Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
    Digits = "01234567890"
    Puncs = ":#"
)

func Keywords() []string {
    return []string {
        "move", "into", "add", "sub", "from",
    }
}

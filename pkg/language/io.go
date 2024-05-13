package language

// IOToken describes a special type of token that denotes a built-in I/O operation.
type IOToken Token

const (
	// IOTokenAsk describes the token that denotes requesting input from stdin.
	IOTokenAsk IOToken = "ask"

	// IOTokenSay describes the token that denotes writing something to stdout.
	IOTokenSay IOToken = "say"
)

// Returns the string representation of the built-in I/O operation token.
func (i IOToken) String() string {
	switch i {
	case IOTokenAsk:
		return "T_IO_ASK"
	case IOTokenSay:
		return "T_IO_SAY"
	}

	// tokens can be extended in the future and we want to make sure there is always a sensible default
	return "T_IO_UNKNOWN"
}

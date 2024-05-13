package language

// SubroutineToken describes a special type of token that denotes a subroutine operation.
type SubroutineToken Token

const (
	// SubroutineTokenCall describes the token that denotes a subroutine is being invoked.
	SubroutineTokenCall SubroutineToken = "call"

	// SubroutineTokenReturn describes the token that denotes a return value sent back when execution completes.
	SubroutineTokenReturn SubroutineToken = "return"

	// SubroutineTokenStart describes the token that denotes the beginning of a subroutine.
	SubroutineTokenStart SubroutineToken = "sub"
)

// Returns the string representation of the subroutine token.
func (s SubroutineToken) String() string {
	switch s {
	case SubroutineTokenCall:
		return "T_SUBROUTINE_CALL"
	case SubroutineTokenReturn:
		return "T_SUBROUTINE_RETURN"
	case SubroutineTokenStart:
		return "T_SUBROUTINE_START"
	}

	// tokens can be extended in the future and we want to make sure there is always a sensible default
	return "T_SUBROUTINE_UNKNOWN"
}

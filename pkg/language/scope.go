package language

// ScopeToken describes a special type of token that denotes a change in scoping or depth.
type ScopeToken Token

const (
	// ScopeTokenEnd describes the token that designates the current scope is ending.
	ScopeTokenEnd ScopeToken = "}"

	// ScopeTokenStart describes the token that designates a new scope is beginning.
	ScopeTokenStart ScopeToken = "{"
)

// Returns the string representation of the scope token.
func (s ScopeToken) String() string {
	switch s {
	case ScopeTokenEnd:
		return "T_SCOPE_END"
	case ScopeTokenStart:
		return "T_SCOPE_START"
	}

	// tokens can be extended in the future and we want to make sure there is always a sensible default
	return "T_SCOPE_UNKNOWN"
}

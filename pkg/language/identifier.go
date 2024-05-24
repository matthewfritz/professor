package language

// IdentifierToken describes a special type of token that denotes a general identifier.
type IdentifierToken Token

const (
	// IdentifierTokenRegex describes the pattern by which identifiers will be evaluated and marked as valid.
	IdentifierTokenRegex = "_?[A-Za-z]+[A-Za-z0-9_]*"
)

// Returns the string representation of the delimiter token.
func (i IdentifierToken) String() string {
	// tokens can be extended in the future and we want to make sure there is always a sensible default
	return "T_IDENTIFIER"
}

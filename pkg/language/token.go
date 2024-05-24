package language

type Tokenized interface {
	String() string
}

// Token is the base type of pretty much every other part of the language.
type Token string

const (
	// TokenUnknown describes a token that could not be resolved by the lexer.
	TokenUnknown Token = ""
)

// Returns the string representation of the token.
func (t Token) String() string {
	switch t {
	case TokenUnknown:
		return "T_TOKEN_UNKNOWN"
	}

	// tokens can be extended in the future and we want to make sure there is always a sensible default
	return "T_TOKEN_UNKNOWN"
}

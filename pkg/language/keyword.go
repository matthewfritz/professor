package language

// KeywordToken describes a special type of token that denotes a general keyword or reserved word.
type KeywordToken Token

const (
	// KeywordTokenBooleanFalse describes the token that denotes a 'false' value as a boolean.
	KeywordTokenBooleanFalse KeywordToken = "FALSE"

	// KeywordTokenBooleanTrue describes the token that denotes a 'true' value as a boolean.
	KeywordTokenBooleanTrue KeywordToken = "TRUE"

	// KeywordTokenNull describes the token that denotes a 'null' value in a variable.
	KeywordTokenNull KeywordToken = "NULL"
)

// Returns the string representation of the keyword token.
func (k KeywordToken) String() string {
	switch k {
	case KeywordTokenBooleanFalse:
		return "T_KEYWORD_BOOLEAN_FALSE"
	case KeywordTokenBooleanTrue:
		return "T_KEYWORD_BOOLEAN_TRUE"
	case KeywordTokenNull:
		return "T_KEYWORD_NULL"
	}

	// tokens can be extended in the future and we want to make sure there is always a sensible default
	return "T_KEYWORD_UNKNOWN"
}

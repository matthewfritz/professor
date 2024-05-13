package language

// DelimiterToken describes a special type of token that denotes a general delimiter.
type DelimiterToken Token

const (
	// DelimiterTokenCharacter describes the token that denotes an enclosed character value.
	DelimiterTokenCharacter DelimiterToken = "'"

	// DelimiterTokenEscape describes the token that denotes the beginning of an escape sequence in a string.
	DelimiterTokenEscape DelimiterToken = "\\"

	// DelimiterTokenString describes the token that denotes an enclosed string value.
	DelimiterTokenString DelimiterToken = "\""
)

// Returns the string representation of the delimiter token.
func (d DelimiterToken) String() string {
	switch d {
	case DelimiterTokenCharacter:
		return "T_DELIMITER_CHARACTER"
	case DelimiterTokenEscape:
		return "T_DELIMITER_ESCAPE"
	case DelimiterTokenString:
		return "T_DELIMITER_STRING"
	}

	// tokens can be extended in the future and we want to make sure there is always a sensible default
	return "T_DELIMITER_UNKNOWN"
}

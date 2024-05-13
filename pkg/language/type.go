package language

// TypeToken describes a special type of token that denotes a built-in type declaration.
type TypeToken Token

const (
	// TypeTokenBoolean describes the token that denotes a boolean declaration.
	TypeTokenBoolean TypeToken = "bool"

	// TypeTokenCharacter describes the token that denotes a single-character declaration.
	TypeTokenCharacter TypeToken = "char"

	// TypeTokenFloat describes the token that denotes a floating-point number declaration.
	TypeTokenFloat TypeToken = "float"

	// TypeTokenInteger describes the token that denotes an integer declaration.
	TypeTokenInteger TypeToken = "int"

	// TypeTokenString describes the token that denotes a string declaration.
	TypeTokenString TypeToken = "string"

	// TypeTokenVoid describes the token that denotes a void declaration.
	TypeTokenVoid TypeToken = "void"
)

// Returns the string representation of the type declaration token.
func (t TypeToken) String() string {
	switch t {
	case TypeTokenBoolean:
		return "T_TYPE_BOOLEAN"
	case TypeTokenCharacter:
		return "T_TYPE_CHARACTER"
	case TypeTokenFloat:
		return "T_TYPE_FLOAT"
	case TypeTokenInteger:
		return "T_TYPE_INTEGER"
	case TypeTokenString:
		return "T_TYPE_STRING"
	case TypeTokenVoid:
		return "T_TYPE_VOID"
	}

	// tokens can be extended in the future and we want to make sure there is always a sensible default
	return "T_TYPE_UNKNOWN"
}

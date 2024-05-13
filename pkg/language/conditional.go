package language

// ConditionalToken describes a special type of token that denotes conditional operations.
type ConditionalToken Token

const (
	// ConditionalTokenFinalTest describes the token that opens the final test statement.
	ConditionalTokenFinalTest ConditionalToken = "else"

	// ConditionalTokenNextTest describes the token that opens the following test statement.
	ConditionalTokenNextTest ConditionalToken = "elseif"

	// ConditionalTokenStart describes the token that begins a conditional statement.
	ConditionalTokenStart ConditionalToken = "if"
)

// Returns the string representation of the conditional token.
func (c ConditionalToken) String() string {
	switch c {
	case ConditionalTokenFinalTest:
		return "T_CONDITIONAL_FINAL_TEST"
	case ConditionalTokenNextTest:
		return "T_CONDITIONAL_NEXT_TEST"
	case ConditionalTokenStart:
		return "T_CONDITIONAL_START"
	}

	// tokens can be extended in the future and we want to make sure there is always a sensible default
	return "T_CONDITIONAL_UNKNOWN"
}

package language

// LoopToken describes a special type of token that denotes looping operations.
type LoopToken Token

const (
	// LoopTokenRunOneOrMoreTimes describes the token that opens a loop where execution happens at least once.
	LoopTokenRunOneOrMoreTimes LoopToken = "do"

	// LoopTokenRunWithVariableControl describes the token that opens a loop where execution happens zero or more times
	// and a specific variable is tracked during iteration.
	LoopTokenRunWithVariableControl LoopToken = "for"

	// LoopTokenRunZeroOrMoreTimes describes the token that opens a loop where execution happens zero or more times.
	LoopTokenRunZeroOrMoreTimes LoopToken = "while"
)

// Returns the string representation of the loop token.
func (l LoopToken) String() string {
	switch l {
	case LoopTokenRunOneOrMoreTimes:
		return "T_LOOP_RUN_ONE_OR_MORE_TIMES"
	case LoopTokenRunWithVariableControl:
		return "T_LOOP_RUN_WITH_VARIABLE_CONTROL"
	case LoopTokenRunZeroOrMoreTimes:
		return "T_LOOP_RUN_ZERO_OR_MORE_TIMES"
	}

	// tokens can be extended in the future and we want to make sure there is always a sensible default
	return "T_LOOP_UNKNOWN"
}

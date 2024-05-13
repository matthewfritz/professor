package language

// OperatorToken describes a special type of token that denotes a built-in operator.
type OperatorToken Token

const (
	// OperatorTokenAssignment describes the token that denotes an assignment operation.
	OperatorTokenAssignment OperatorToken = "="

	// OperatorTokenEquality describes the token that denotes an equality check.
	OperatorTokenEquality OperatorToken = "=="

	// OperatorTokenEqualityAndSameType describes the token that denotes an equality + type check.
	OperatorTokenEqualityAndSameType OperatorToken = "==="

	// OperatorTokenGreaterThan describes the token that denotes a GT check.
	OperatorTokenGreaterThan OperatorToken = ">"

	// OperatorTokenGreaterThanOrEqual describes the token that denotes a GTE check.
	OperatorTokenGreaterThanOrEqual OperatorToken = ">="

	// OperatorTokenLessThan describes the token that denotes an LT check.
	OperatorTokenLessThan OperatorToken = "<"

	// OperatorTokenLessThanOrEqual describes the token that denotes an LTE check.
	OperatorTokenLessThanOrEqual OperatorToken = "<="

	// OperatorTokenLogicalAnd describes the token that denotes a logical AND.
	OperatorTokenLogicalAnd OperatorToken = "and"

	// OperatorTokenLogicalOr describes the token that denotes a logical OR.
	OperatorTokenLogicalOr OperatorToken = "or"
)

// Returns the string representation of the operator declaration token.
func (o OperatorToken) String() string {
	switch o {
	case OperatorTokenAssignment:
		return "T_OPERATOR_ASSIGNMENT"
	case OperatorTokenEquality:
		return "T_OPERATOR_EQUALITY"
	case OperatorTokenEqualityAndSameType:
		return "T_OPERATOR_EQUALITY_SAME_TYPE"
	case OperatorTokenGreaterThan:
		return "T_OPERATOR_GREATER_THAN"
	case OperatorTokenGreaterThanOrEqual:
		return "T_OPERATOR_GREATER_THAN_OR_EQUAL"
	case OperatorTokenLessThan:
		return "T_OPERATOR_LESS_THAN"
	case OperatorTokenLessThanOrEqual:
		return "T_OPERATOR_LESS_THAN_OR_EQUAL"
	case OperatorTokenLogicalAnd:
		return "T_OPERATOR_LOGICAL_AND"
	case OperatorTokenLogicalOr:
		return "T_OPERATOR_LOGICAL_OR"
	}

	// tokens can be extended in the future and we want to make sure there is always a sensible default
	return "T_OPERATOR_UNKNOWN"
}

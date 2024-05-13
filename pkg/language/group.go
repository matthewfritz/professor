package language

// GroupToken describes a special type of token that denotes a change in grouping but not scope.
type GroupToken Token

const (
	// GroupTokenEnd describes the token that designates the current group is ending.
	GroupTokenEnd GroupToken = ")"

	// GroupTokenStart describes the token that designates a new group is starting.
	GroupTokenStart GroupToken = "("
)

// Returns the string representation of the group token.
func (g GroupToken) String() string {
	switch g {
	case GroupTokenEnd:
		return "T_GROUP_END"
	case GroupTokenStart:
		return "T_GROUP_START"
	}

	// tokens can be extended in the future and we want to make sure there is always a sensible default
	return "T_GROUP_UNKNOWN"
}

package core

// Group is an Artifact that provides a visual mechanism to group elements
// of a diagram informally.
// The grouping is tied to the CategoryValue supporting element.
// That is, a Group is a visual depiction of a single CategoryValue.
type Group struct {
	*Artifact
	// CategoryValueRef specifies the CategoryValue that the Group represents.
	CategoryValue *CategoryValue
}

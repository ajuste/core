package core

type Artifact struct {
	*BaseElement
}

type AssociationDirection string

const (
	AssociationDirectionNone = AssociationDirection("None")
	AssociationDirectionOne  = AssociationDirection("One")
	AssociationDirectionBoth = AssociationDirection("Both")
)

// Association  is used to associate information and Artifacts with Flow Objects.
// Text and graphical non-Flow Objects can be associated with the Flow Objects and Flow.
// An Association is also used to show the Activity used for compensation
type Association struct {
	*Artifact
	// AssociationDirection defines whether or not the Association shows any
	// directionality with an arrowhead. The default is None (no arrowhead).
	// A value of One means that the arrowhead SHALL be at the Target Object.
	// A value of Both means that there SHALL be an arrowhead at both ends of the Association line.
	AssociationDirection AssociationDirection
	// SourceRef the Association is connecting from.
	Source *BaseElement
	// TargetRef the Association is connecting to.
	Target *BaseElement
}

// Category which have user-defined semantics, can be used for documentation or analysis purpose
type Category struct {
	// Name descriptive name of the element.
	Name string
	// CategoryValue specifies one or more values of the Category.
	CategoryValue []*CategoryValue
}

// CategoryValue for a category
type CategoryValue struct {
	*BaseElement
	// Value of the CategoryValue element.
	Value string
	// Category specifies the Category representing the Category as such and contains the CategoryValue
	Category []*Category
	// CategorizedFlowElements identifies all of the elements (e.g., Events, Activities, Gateways, and Artifacts)
	// that are within the boundaries of the Group.
	CategorizedFlowElements []*FlowElement
}

// TextAnnotation are a mechanism for a modeler to provide additional information for the reader of a BPMN Diagram.
type TextAnnotation struct {
	*Artifact
	// Text that the modeler wishes to communicate to the reader of the Diagram.
	Text string
	// TextFormat of the text. It MUST follow the mime- type format. The default is "text/plain."
	TextFormat string
}

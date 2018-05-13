package core

// ItemKind specifies the nature of an item which can be a physical or an information item.
type ItemKind int

const (
	// Physical item kind
	Physical = ItemKind(0)

	// Information item kind
	Information = ItemKind(1)
)

// ItemDefinition BPMN elements, such as DataObjects and Messages,
// represent items that are manipulated, transferred, transformed, or
// stored during Process flows. These items can be either physical items,
// such as the mechanical part of a vehicle, or information items such the catalog of the mechanical parts of a vehicle.
type ItemDefinition struct {
	// ItemKind defines nature of the Item.
	ItemKind ItemKind

	// Structure is the concrete data structure to be used.
	Structure *Element

	// IsCollection set to true when item difinition is for a collection
	IsCollection bool

	// Import identifies the location of the data structure and its format.
	// If the importType attribute is left unspecified,
	// the typeLanguage specified in the Definitions that contains this ItemDefinition is assumed.
	Import *Import

	Message *Message
}

// Message represents the content of a communication between two Participants.
// In BPMN 2.0, a Message is a graphical decorator (it was a supporting element in BPMN 1.2).
// An ItemDefinition is used to specify the Message structure.
type Message struct {
	*BaseElement
	Name string
	Item *ItemDefinition
}

type Resource struct {
	Name               string
	ResourceParameters []*ResourceParameter
}

type ResourceParameter struct {
	Name       string
	IsRequired bool
}

type Expression struct {
	*BaseElement
}

// CallableElement is the abstract super class of all Activities that have been
// defined outside of a Process or Choreography but which can be called (or reused),
// by a Call Activity, from within a Process or Choreography.
// It MAY reference Interfaces that define the service operations that it provides.
type CallableElement struct {

	// IOSpecification defines the inputs and outputs and the InputSets and OutputSets for the Activity.
	IOSpecification *InputOutputSpecification

	// IOBinding defines a combination of one InputSet and one OutputSet in order to bind this to an operation defined in an interface.
	IOBinding *InputOutputBinding

	// SupportedInterfaces interfaces describing the external behavior provided by this element.
	SupportedInterfaces *Interface

	// Name descriptive name of the element.
	Name string
}

// Error represents the content of an Error Event or the Fault of a failed Operation.
// An ItemDefinition is used to specify the structure of the Error.
type Error struct {
	*RootElement
	Name string
	// ErrorCode For an End Event:
	// If the result is an Error, then the errorCode MUST be supplied (if the processType attribute of the Process is set to execut- able) This “throws” the Error.
	// For an Intermediate Event within normal flow:
	// If the trigger is an Error, then the errorCode MUST be entered
	// (if the processType attribute of the Process is set to execut-
	// able). This “throws” the Error.
	// For an Intermediate Event attached to the boundary of an Activity:
	// If the trigger is an Error, then the errorCode MAY be entered. This Event “catches” the Error. If there is no errorCode, then any error SHALL trigger the Event. If there is an errorCode, then only an Error that matches the errorCode SHALL trigger the Event.
	ErrorCode string
	IsClosed  bool
	// Structure defines payload of the error
	Structure *ItemDefinition
}

// CorrelationKey represents a composite key out of one or many CorrelationProperties that essentially specify extraction Expressions atop Messages.
type CorrelationKey struct {
	*Artifact
	Name string
	// CorrelationPropertyRef representing the partial keys of this CorrelationKey.
	CorrelationPropertyRef []*CorrelationProperty
	ConversationNode       *ConversationNode
}

type CorrelationSubscription struct {
	*BaseElement
	Name                                   string
	Type                                   *ItemDefinition
	CorrelationPropertyRetrievalExpression []*CorrelationPropertyRetrievalExpression
}

type CorrelationProperty struct {
	*BaseElement
	Name string
	Type *ItemDefinition
	// CorrelationPropertyRetrievalExpressions for this CorrelationProperty,
	// representing the associations of FormalExpressions (extraction paths) to specific Messages occurring in this Conversation.
	CorrelationPropertyRetrievalExpressions []*CorrelationPropertyRetrievalExpression
}

type CorrelationPropertyRetrievalExpression struct {
	*BaseElement
	CorrelationProperty *CorrelationProperty
	// MessagePath is the FormalExpression that defines how to extract a CorrelationProperty from the Message payload.
	MessagePath *FormalExpression
	// Message the FormalExpression extracts the CorrelationProperty from.
	Message *Message
}

type CorrelationPropertyBinding struct {
	// DataPath FormalExpression that defines the extraction rule atop the Process context.
	DataPath            string
	CorrelationProperty *CorrelationProperty
}

// FormalExpression specify an executable Expression using a specified Expression language.
// A natural-language description of the Expression can also be specified, in addition to the formal specification.
type FormalExpression struct {
	Language        string
	Body            *Element
	CorrelationSet  *CorrelationPropertyRetrievalExpression
	EvaluatesToType *ItemDefinition
}

// FlowElement super class for all elements that can appear in a Process flow, which are FlowNodes
type FlowElement struct {
	*BaseElement
	Name string
}

// FlowElementsContainer abstract super class for BPMN diagrams (or views) and
// defines the superset of elements that are contained in those diagrams.
// Basically, a FlowElementsContainer contains FlowElements,
// which are Events, Gateways, Sequence Flows, Activities, and Choreography Activities.
type FlowElementsContainer struct {
	*BaseElement
	// FlowElements specifies the particular flow elements contained in a FlowElementContainer.
	// Flow elements are Events, Gateways, Sequence Flows, Activities, Data Objects,
	// Data Associations, and Choreography Activities.
	//	Note that:
	//	• ChoreographyActivitiesMUSTNOTbeincludedasaflowElementfora Process.
	//	• Activities, Data Associations, and Data Objects MUST NOT be included as a flowElement for a Choreography.

	FlowElements []*FlowElement

	// LaneSets list of LaneSets used in the FlowElementsContainer LaneSets are not used for Choreographies or Sub-Choreographies.
	LaneSets []*LaneSet
}

type FlowNode struct {
	*BaseElement
	Name string
}

// SequenceFlow is used to show the order of Flow Elements in a Process or a Choreography.
type SequenceFlow struct {
	*FlowElement

	// IsImmediate An optional boolean value specifying whether Activities or Choreography Activities not in the model containing the Sequence Flow can occur between the elements connected by the Sequence Flow. If the value is true, they MAY NOT occur. If the value is false, they MAY occur. Also see the isClosed attribute on Process, Choreography, and Collaboration. When the attribute has no value, the default semantics depends on the kind of model containing Sequence Flows:
	// • For non-executable Processes (public Processes and non-executable private Processes) and Choreographies no value has the same semantics as if the value were false.
	// • For an executable Processes no value has the same semantics as if the value were true.
	// • For executable Processes, the attribute MUST NOT be false.
	IsImmediate bool

	// Source is the  FlowNode that the Sequence Flow is connecting from.
	// For a Process: Of the types of FlowNode, only Activities, Gateways, and Events can be the source. However, Activities that are Event Sub-Processes are not allowed to be a source.
	// For a Choreography: Of the types of FlowNode, only Choreography Activities, Gateways, and Events can be the source.
	Source *FlowNode

	// Target is the FlowNode that the Sequence Flow is connecting to.
	// For a Process: Of the types of FlowNode, only Activities, Gateways, and Events can be the target. However, Activities that are Event Sub-Processes are not allowed to be a target.
	// For a Choreography: Of the types of FlowNode, only Choreography Activities, Gateways, and Events can be the target.
	Target *FlowNode

	// ConditionExpression optional boolean Expression that acts as a gating condition. A token will only be placed on this Sequence Flow if this conditionExpression evaluates to true.
	ConditionExpression *Expression
}

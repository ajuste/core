package core

type EventDefinition struct {
}

type Event struct {
	*FlowNode
	Properties []*Property
}

type ThrowEvent struct {
	*Event
	DataInputs []*DataInput
}

type ImplicitThrowEvent struct {
	*ThrowEvent
}

type IntermediateThrowEvent struct {
	*ThrowEvent
}

type EndEvent struct {
	*ThrowEvent
}

type CatchEvent struct {
	*Event
	DataOutputs []*DataOutput
}

type StartEvent struct {
	*CatchEvent
}

type IntermediateCatchEvent struct {
	*CatchEvent
}

type BoundaryEvent struct {
	*CatchEvent
	CancelActivity bool
	AttachedTo     *Activity
}

// Escalation identifies a business situation that a Process might need to react to.
// An ItemDefinition is used to specify the structure of the Escalation.
type Escalation struct {
	Name string

	EscalationCode string

	// Structure define the “payload” of the Escalation.
	Structure *ItemDefinition
}

type MessageEventDefinition struct {
	Message *Message
}

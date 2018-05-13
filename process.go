package core

// ProcessType is level of abstraction modeled by a Process
type ProcessType int32

const (
	// ProcessTypeNone a public Process shows only those flow elements that are relevant to external consumers. Internal details are not modeled
	ProcessTypeNone = ProcessType(0)
	// ProcessTypePrivate is one that is internal to a specific organization.
	ProcessTypePrivate = ProcessType(1)
	// ProcessTypePublic the processType is “none,” meaning undefined.
	ProcessTypePublic = ProcessType(2)
)

// Process describes a sequence or flow of Activities in an organization with
// the objective of carrying out work. In BPMN a Process is depicted as a
// graph of Flow Elements, which are a set of Activities, Events, Gateways,
// and Sequence Flows that define finite execution semantics.
// Processes can be defined at any level from enterprise-wide Processes to
// Processes performed by a single person. Low-level Processes can be grouped
// together to achieve a common business goal.
type Process struct {
	*FlowElementsContainer
	*CallableElement
	// IsClosed A boolean value specifying whether interactions,
	// such as sending and receiving Messages and Events,
	// not modeled in the Process can occur when the Process is executed or performed.
	// If the value is true, they MAY NOT occur. If the value is false, they MAY occur.
	IsClosed bool

	// IsExecutable optional Boolean value specifying whether the Process is execut- able.
	// An executable Process is a private Process that has been modeled for the purpose of being executed
	IsExecutable bool

	// Artifacts provides the list of Artifacts that are contained within the Process.
	Artifacts []*Artifact

	// CorrelationSubscriptions are used to correlate incoming Messages against data in the Process context. A Process MAY contain several correlationSubscriptions
	CorrelationSubscriptions   []*CorrelationSubscription
	CorrelationPropertyBinding []*CorrelationPropertyBinding
	Properties                 []*Property

	// Supports modelers can declare that they intend all executions or
	// performances of one Process to also be valid for another Process.
	// This means they expect all the executions or performances of the
	// first Processes to also follow the steps laid out in the second Process.
	Supports []*Process

	// DefinitionalCollaboration for Processes that interact with other Participants,
	// a definitional Collaboration can be referenced by the Process.
	// The definitional Collaboration specifies the Participants the Process
	// interacts with, and more specifically, which individual service,
	// Send or Receive Task, or Message Event, is connected to which Participant
	// through Message Flows. The definitional Collaboration need not be dis- played.
	// Additionally, the definitional Collaboration can be used to include Con- versation information within a Process.
	DefinitionalCollaboration []*Collaboration

	// Auditing provides a hook for specifying audit related properties.
	// Monitoring provides a hook for specifying monitoring related properties.
	// Resources Defines the resource that will perform or will be responsible for the Process. The resource, e.g., a performer, can be specified in the form of a specific individual, a group, an organization role or position, or an organization.
}

// Performer resource that will perform or will be responsible for an Activity
type Performer struct {
	*ResourceRole
}

type LaneSet struct {
	*BaseElement
	Name string
}

// GlobalTask is a reusable, atomic Task definition that can be called from within any Process by a Call Activity.
type GlobalTask struct {
}

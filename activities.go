package core

// Activity is work that is performed within a Business Process.
// An Activity can be atomic or non-atomic (compound).
// The types of Activities that are a part of a Process are: Task, Sub-Process,
// and Call Activity, which allows the inclusion of re-usable Tasks and
// Processes in the diagram.
type Activity struct {
	*FlowNode

	// IsForCompensation If false, then this Activity executes as a result of
	// normal execution flow. If true, this Activity is only activated when a
	// Compensation Event is detected and initiated under Compensation Event visibility scop
	IsForCompensation bool

	// StartQuantity number of tokens that MUST arrive before the Activity can begin.
	// The default value is 1. The value MUST NOT be less than 1.
	StartQuantity int

	// CompletionQuantity defines the number of tokens that MUST be generated from the Activity.
	// This number of tokens will be sent done any outgoing Sequence Flow (assuming any Sequence Flow conditions are satis- fied)
	CompletionQuantity int

	// BoundaryEvents references the Intermediate Events that are attached to the boundary of the Activity.
	BoundaryEvents []*BoundaryEvent

	// IOSpecification defines the inputs and outputs and the InputSets and OutputSets for the Activit
	IOSpecification *InputOutputSpecification

	// DataOutputAssociations an optional reference to the DataOutputAssociations.
	DataOutputAssociations []*DataOutputAssociation

	// DataInputAssociations an optional reference to the DataInputAssociations
	// A DataInputAssociation defines how the DataInput of the Activity’s
	// InputOutputSpecification will be populated.
	DataInputAssociations []*DataInputAssociation

	// Properties modeler-defined properties MAY be added to an Activity.
	// These properties are contained within the Activity.
	Properties []*Property

	// Default is the Sequence Flow that will receive a token when
	// none of the conditionExpressions on other outgoing Sequence Flows evaluate to true.
	// The default Sequence Flow should not have a conditionExpression.
	// Any such Expression SHALL be ignored.
	Default *SequenceFlow

	// LoopCharacteristics an Activity MAY be performed once or MAY be repeated.
	// If repeated, the Activity MUST have loopCharacteristics that define the
	// repetition criteria (if the isExecutable attribute of the Process is set to true).
	LoopCharacteristics *LoopCharacteristics

	// Resources defines the resource that will perform or will be responsible for the Activity.
	// The resource, e.g., a performer, can be specified in the form of a specific individual,
	// a group, an organization role or position, or an organization.
	Resources []*ResourceRole
}

// CallActivity identifies a point in the Process where a global Process or
// a Global Task is used. The Call Activity acts as a ‘wrapper’ for the
// invocation of a global Process or Global Task within the execution.
// The activation of a call Activity results in the transfer of control to the called global Process or Global Task.
type CallActivity struct {
	*Activity

	// CalledElement element to be called, which will be either a Process or a GlobalTask.
	// Other CallableElements, such as Choreography, GlobalChoreographyTask,
	// Conversation, and GlobalCommunication MUST NOT be called by the Call Conversation element.
	CalledElement *CallableElement
}

// Task is an atomic Activity within a Process flow.
// A Task is used when the work in the Process cannot be broken down
// to a finer level of detail. Generally, an end-user and/or applications
// are used to perform the Task when it is executed.
type Task struct {
	*Activity
}

// ReceiveTask is a simple Task that is designed to wait for a Message to
// rrive from an external Participant (relative to the Process).
// Once the Message has been received, the Task is completed.
type ReceiveTask struct {
	*Task

	// Implementation specifies the technology that will be used to send
	// and receive the Messages. Valid values are "##unspecified" for
	// leaving the implementation technology open, "##WebService" for the Web
	// service technology or a URI identifying any other technology or
	// coordination protocol. A Web service is the default technology.
	Implementation string

	// Instantiate Receive Tasks can be defined as the instantiation
	// mechanism for the Process with the instantiate attribute.
	// This attribute MAY be set to true if the Task is the first Activity (i.e.,
	// there are no incoming Sequence Flows).
	// Multiple Tasks MAY have this attribute set to true.
	Instantiate bool

	// Message for the messageRef attribute MAY be entered.
	// This indicates that the Message will be received by the Task.
	// The Message in this context is equivalent to an in-only message pattern (Web service).
	// One (1) or more corresponding incoming Message Flows MAY be shown on the diagram.
	// However, the display of the Message Flows is NOT REQUIRED.
	// The Message is applied to all incoming Message Flows,
	// but can arrive for only one (1) of the incoming Message Flows for a single instance of the Task.
	Message *Message

	// Operation specifies the operation that is invoked by the Service Task.
	Operation *Operation
}

// ServiceTask is a Task that uses some sort of service,
// which could be a Web service or an automated application.
type ServiceTask struct {
	*Task

	// Implementation specifies the technology that will be used to send
	// and receive the Messages. Valid values are "##unspecified" for
	// leaving the implementation technology open, "##WebService" for the Web
	// service technology or a URI identifying any other technology or
	// coordination protocol. A Web service is the default technology.
	Implementation string

	// Operation specifies the operation that is invoked by the Service Task.
	Operation *Operation
}

// SendTask is a simple Task that is designed to send a Message to an
// external Participant (relative to the Process).
// Once the Message has been sent, the Task is completed.
type SendTask struct {
	*Task

	// Implementation specifies the technology that will be used to send
	// and receive the Messages. Valid values are "##unspecified" for
	// leaving the implementation technology open, "##WebService" for the Web
	// service technology or a URI identifying any other technology or
	// coordination protocol. A Web service is the default technology.
	Implementation string

	// Message indicates that the Message will be sent by the Task.
	// The Message in this context is equivalent to an out-only message
	// pattern (Web service). One or more corresponding outgoing Message
	// Flows MAY be shown on the diagram. However, the display of the Message
	// Flows is NOT REQUIRED. The Message is applied to all outgoing Message
	// Flows and the Message will be sent down all outgoing Message Flows at
	// the completion of a single instance of the Task.
	Message *Message

	// Operation that is invoked by the Send Task.
	Operation *Operation
}

// ScriptTask is executed by a business process engine.
// The modeler or implementer defines a script in a language that the
// engine can interpret. When the Task is ready to start,
// the engine will execute the script. When the script is completed,
// the Task will also be completed.
type ScriptTask struct {
	*Task

	// ScriptFormat defines the format of the script.
	// This attribute value MUST be specified with a mime-type format.
	// And it MUST be specified if a script is provided.
	ScriptFormat string

	// Script to be run
	Script string
}

// BusinessRuleTask provides a mechanism for the Process to provide
// input to a Business Rules Engine and to get the output of calculations
// that the Business Rules Engine might provide.
// The InputOutputSpecification of the Task will allow
// the Process to send data to and receive data from the Business Rules Engine.
type BusinessRuleTask struct {
	*Task

	// Implementation specifies the technology that will be used to send
	// and receive the Messages. Valid values are "##unspecified" for
	// leaving the implementation technology open, "##WebService" for the Web
	// service technology or a URI identifying any other technology or
	// coordination protocol. A Web service is the default technology.
	Implementation string
}

// SubProcess is an Activity whose internal details have been modeled using
// Activities, Gateways, Events, and Sequence Flows.
// A Sub-Process is a graphical object within a Process,
// but it also can be “opened up” to show a lower-level Process.
// Sub-Processes define a contextual scope that can be used for attribute visibility,
// transactional scope, for the handling of exceptions, of Events, or for compensation
type SubProcess struct {
	*FlowElementsContainer
	*Activity
	Artifacts []*Artifact

	// TriggeredByEvent a flag that identifies whether this Sub-Process is an Event Sub-Process.
	// • If false, then this Sub-Process is a normal Sub-Process.
	// • If true, then this Sub-Process is an Event Sub-Process and is subject to additional constraints.
	TriggeredByEvent bool
}

// LoopCharacteristics activities MAY be repeated sequentially, essentially behaving like a loop. The presence of LoopCharacteristics signifies that the Activity has looping behavior. LoopCharacteristics is an abstract class. Concrete subclasses define specific kinds of looping behavior.
type LoopCharacteristics struct {
	// Left empty intentionally
}

// StandardLoopCharacteristics defines looping behavior based on a boolean condition.
// The Activity will loop as long as the boolean condition is true.
// The condition is evaluated for every loop iteration, and MAY be evaluated at the beginning or at the end of the iteration.
// In addition, a numeric cap can be optionally specified.
type StandardLoopCharacteristics struct {
	*LoopCharacteristics

	// TestBefore flag that controls whether the loop condition is evaluated
	// at the beginning (testBefore = true) or at the end (testBefore = false) of the loop iteration.
	TestBefore bool

	// LoopMaximum serves as a cap on the number of iterations.
	LoopMaximum *Expression

	// LoopCondition is a boolean Expression that controls the loop.
	// The Activity will only loop as long as this condition is true.
	// The looping behavior MAY be underspecified, meaning that the modeler can
	// simply document the condition, in which case the loop cannot be formally executed.
	LoopCondition *Expression
}

type MultiInstanceBehavior string

const (

	// MultiInstanceBehaviorNone the EventDefinition which is associated through the noneEvent association will be thrown for each instance completing.
	MultiInstanceBehaviorNone = MultiInstanceBehavior("None")

	// MultiInstanceBehaviorOne the EventDefinition referenced through the oneEvent association will be thrown upon the first instance completing.
	MultiInstanceBehaviorOne = MultiInstanceBehavior("One")

	// MultiInstanceBehaviorAll no Event is ever thrown; a token is produced after completion of all instances.
	MultiInstanceBehaviorAll = MultiInstanceBehavior("All")

	// MultiInstanceBehaviorComplex the complexBehaviorDefinitions are consulted to determine if and which Events to throw.
	MultiInstanceBehaviorComplex = MultiInstanceBehavior("Complex")
)

// MultiInstanceLoopCharacteristics class allows for creation of a desired number of Activity instances.
// The instances MAY execute in parallel or MAY be sequential.
// Either an Expression is used to specify or calculate the desired number of
// instances or a data driven setup can be used. In that case a data input can be specified,
// which is able to handle a collection of data. The number of items in the collection determines
// the number of Activity instances. This data input can be produced by an input Data Association.
// The modeler can also configure this loop to control the tokens produced.
type MultiInstanceLoopCharacteristics struct {
	*LoopCharacteristics

	// IsSequential is a flag that controls whether the Activity instances will execute sequentially or in parallel.
	IsSequential bool

	// Behavior acts as a shortcut for specifying when events SHALL be thrown from an Activity instance that is about to complete.
	Behavior MultiInstanceBehavior

	// LoopCardinality A numeric Expression that controls the number of Activity instances that will be created. This Expression MUST evaluate to an integer.
	// This MAY be underspecified, meaning that the modeler MAY simply doc- ument the condition. In such a case the loop cannot be formally executed.
	// In order to initialize a valid multi-instance, either the loopCardinality Expression or the loopDataInput MUST be specified.
	LoopCardinality *Expression

	// CompletionCondition defines a boolean Expression that when evaluated to true, cancels the remaining Activity instances and produces a token.
	CompletionCondition *Expression

	// InputDataItem a Data Input, representing for every Activity instance
	// the single item of the collection stored in the loopDataInput.
	// This Data Input can be the source of DataInputAssociation to a data
	// input of the Activity’s InputOutputSpecification.
	// The type of this Data Input MUST the scalar of the type defined for the loopDataInput.
	InputDataItem *DataInput

	// OutputDataItem a Data Output, representing for every Activity instance
	// the single item of the collection stored in the loopDataOutput.
	// This Data Output can be the target of DataOutputAssociation to a
	// data output of the Activity’s InputOutputSpecification.
	// The type of this Data Output MUST the scalar of the type defined for the loopDataOutput.
	OutputDataItem *DataOutput

	// LoopDataOutput is a ItemAwareElement specifies the collection of data, which will be produced by the multi-instance.
	// For Tasks it is a reference to a Data Output which is part of the Activity’s InputOutputSpecification.
	// For Sub-Processes it is a reference to a collection-valued Data Object in the context that is visible to the Sub-Processes.
	LoopDataOutput *ItemAwareElement

	// LoopDataInput is a ItemAwareElement is used to determine the number of Activity instances,
	// one Activity instance per item in the collection of data stored in that ItemAwareElement element.
	// For Tasks it is a reference to a Data Input which is part of the Activity’s InputOutputSpecification.
	// For Sub-Processes it is a reference to a collection-valued Data Object in the context that is visible to the Sub-Processes.
	// In order to initialize a valid multi-instance, either the loopCardinality Expression or the loopDataInput MUST be specified.
	LoopDataInput *ItemAwareElement

	// OneBehaviourElement the EventDefinition which is thrown when behavior is set to one and the first internal Activity instance has completed.
	OneBehaviourElement *EventDefinition

	// ComplexBahaviorDefinition controls when and
	// which Events are thrown in case behavior is set to complex.
	ComplexBahaviorDefinition *ComplexBahaviorDefinition

	// NoneBehaviorEvent the EventDefinition which is thrown when
	// the behavior is set to none and an internal Activity instance has completed.
	NoneBehaviorEvent *EventDefinition
}

type MultiInstanceActivityInstance struct {

	// LoopCounter attribute is provided for each generated (inner)
	// instance of the Activity. It contains the sequence number of the
	// generated instance, i.e., if this value of some instance in n, the instance is the n-th instance that was generated.
	LoopCounter int

	// NumberOfInstances attribute is provided for the outer instance of the
	// Multi-Instance Activity only. This attribute contains the total number of
	// inner instances created for the Multi-Instance Activity
	NumberOfInstances int

	// NumberOfActiveInstances attribute is provided for the outer instance
	// of the Multi-Instance Activity only. This attribute contains the number of
	// currently active inner instances for the Multi-Instance Activity.
	// In case of a sequential Multi-Instance Activity,
	// this value can’t be greater than 1. For parallel Multi-Instance Activities,
	// this value can’t be greater than the value contained in numberOfInstance
	NumberOfActiveInstances int

	// NumberOfCompletedInstances attribute is provided for the outer instance
	// of the Multi-Instance Activity only. This attribute contains the
	// number of already com- pleted inner instances for the Multi-Instance Activity.
	NumberOfCompletedInstances int

	// NumberOfTerminatedInstances attribute is provided for the outer
	// instance of the Multi-Instance Activity only. This attribute contains
	// the number of terminated inner instances for the Multi-Instance Activity.
	// The sum of numberOfTerminatedInstances, numberOfCompletedInstances,
	// and numberOfActiveInstances always sums up to numberOfInstances.
	NumberOfTerminatedInstances int
}

// ComplexBahaviorDefinition controls when and which Events
// are thrown in case behavior of the Multi-Instance Activity is set to complex.
type ComplexBahaviorDefinition struct {

	// Condition defines a boolean Expression that when evaluated to true, cancels the remaining Activity instances and produces a token.
	Condition *FormalExpression

	// Event if the condition is true, this identifies the Event that will be thrown (to be caught by a boundary Event on the Multi-Instance Activity).
	Event *ImplicitThrowEvent
}

// ResourceRole element inherits the attributes and model associations of BaseElement.
type ResourceRole struct {
	Name string

	// ResourceAssignmentExpression defines the Expression used for the Resource assignment.
	// Should not be specified when a resourceRef is provided.
	ResourceAssignmentExpression *ResourceAssignmentExpression

	// Resource that is associated with Activity. Should not be specified when resourceAssignmentExpression is provided.
	Resource *Resource

	// ResourceParameterBindings defines the Parameter bindings used for the Resource assignment.
	// Is only applicable if a resource is specified.
	ResourceParameterBindings []*ResourceParameterBinding
}

// ResourceAssignmentExpression must return Resource entity related data types,
// like Users or Groups assigned to Activity.
// Different Expressions can return multiple Resources.
// All of them are assigned to the respective subclass of the ResourceRole element,
// for example as potential owners. The semantics is defined by the subclass.
type ResourceAssignmentExpression struct {
	// Expression used at runtime to assign resource(s) to a ResourceRole element
	Expression Expression
}

type ResourceParameterBinding struct {
	// Expression evaluates the value used to bind the ResourceParameter.
	Expression Expression
	Parameter  ResourceParameter
}

// AdHocOrdering for AdHoc sub processes
type AdHocOrdering string

const (
	// AdHocOrderingParallel all the Activities of the Sub-Process can be performed in parallel.
	AdHocOrderingParallel = AdHocOrdering("Parallel")

	// AdHocOrderingSequential only one Activity can be performed at a time.
	AdHocOrderingSequential = AdHocOrdering("Sequential")
)

// AdHocSubProcess is a specialized type of Sub-Process that is a
// group of Activities that have no REQUIRED sequence relationships.
// A set of Activities can be defined for the Process,
// but the sequence and number of performances for the Activities is
// determined by the performers of the Activities.
type AdHocSubProcess struct {
	*SubProcess

	// Ordering defines if the Activities within the Process can be performed in parallel or MUST be performed sequentially.
	// The default setting is parallel and the setting of sequential is a
	// restriction on the performance that can be needed due to shared resources
	Ordering AdHocOrdering

	// CancelRemainingInstances is used only if ordering is parallel.
	// It determines whether running instances are cancelled when the completionCondition becomes true.
	CancelRemainingInstances bool

	// CompletionCondition defines the conditions when the Process will end.
	// When the Expression is evaluated to true, the Process will be terminated.
	CompletionCondition Expression
}

// Transaction  is a specialized type of Sub-Process that will have a
// special behavior that is controlled through a transaction protocol (such as WS-Transaction).
// The boundary of the Sub-Process will be double-lined to indicate that it is a Transaction
type Transaction struct {
	*SubProcess
	Protocol string

	// Method is an attribute that defines the Transaction method used to
	// commit or cancel a Transaction. For executable Processes,
	// it SHOULD be set to a technology specific URI, e.g.,
	// http://schemas.xmlsoap.org/ws/2004/10/wsat for WS- AtomicTransaction,
	// or http://docs.oasis-open.org/ws-tx/wsba/2006/ 06/AtomicOutcome for WS-BusinessActivity.
	// For compatibility with BPMN 1.1, it can also be set to "##compensate", "##store", or "##image".
	Method string
}

type GlobalScriptTask struct {
	*GlobalTask

	ScriptLanguage string
	Script         string
}

type GlobalBusinessRuleTask struct {
	*GlobalTask

	Implementation string
}

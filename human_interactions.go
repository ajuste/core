package core

// UserTask is a typical “workflow” Task where a human performer performs
// the Task with the assistance of a software application and is scheduled
// through a task list manager of some sort.
type UserTask struct {
	*Task

	// Implementation specifies the technology that will be used to send
	// and receive the Messages. Valid values are "##unspecified" for
	// leaving the implementation technology open, "##WebService" for the Web
	// service technology or a URI identifying any other technology or
	// coordination protocol. A Web service is the default technology.
	Implementation string

	// Renderings acts as a hook which allows BPMN adopters to specify
	// task rendering attributes by using the BPMN Extension mechanism.
	Renderings []*Rendering
}

// UserTaskInstance represents an instace of UserTask
type UserTaskInstance struct {
	// ActualOwner who picked/claimed the User task and
	// became the actual owner of it.
	// The value is a literal representing the user’s id, email address etc.
	ActualOwner string

	TaskPriority int
}

// ManualTask is a Task that is expected to be performed without the aid of
// any business process execution engine or any application.
// An example of this could be a telephone technician installing a telephone at a customer location
type ManualTask struct {
	*Task
	Implementation string
}

type Rendering struct {
	UserTask *UserTask
}

type HumanPerformer struct {
	*Performer
}

// PotentialOwner are persons who can claim and work on it.
// A potential owner becomes the actual owner of a Task, usually by explicitly claiming it
type PotentialOwner struct {
	*HumanPerformer
}

type GlobalUserTask struct {
	*GlobalTask

	Implementation string
	Renderings     []*Rendering
}

type GlobalManualTask struct {
	*GlobalTask
}

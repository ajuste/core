package core

type DataObject struct {
	IsCollection bool
}

type DataOutput struct {
	Name         string
	IsCollection bool
}

type DataInput struct {
	Name         string
	IsCollection bool
}

type InputOutputSpecification struct {
	InputSets  []*InputSet
	OutputSets []*OutputSet
}

type DataInputAssociation struct {
}

type DataOutputAssociation struct {
}

type Property struct {
}

type InputSet struct {
	Name string
}

type OutputSet struct {
	Name string
}

// InputOutputBinding when a CallableElement is exposed as a Service,
// it has to define one or more InputOutputBinding elements.
// An InputOutputBinding element binds one Input and one Output of the
// InputOutputSpecification to an Operation of a Service Interface.
type InputOutputBinding struct {

	// Operation reference to one specific Operation defined as part of the Interface of the Activity.
	Operation *Operation

	// OutputData a reference to one specific DataOutput defined as part of the InputOutputSpecification of the Activity.
	OutputData *DataOutput

	// InputData a reference to one specific DataInput defined as part of the InputOutputSpecification of the Activity.
	InputData *DataInput
}

type ItemAwareElement struct {
}

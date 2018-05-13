package core

// BaseElement is the abstract super class for most BPMN elements.
// It provides the attributes id and documentation, which other elements will inherit.
type BaseElement struct {
	// ID This attribute is used to uniquely identify BPMN elements.
	// The id is REQUIRED if this element is referenced or
	// intended to be referenced by something else.
	// If the element is not currently referenced and is never intended to be referenced, the id MAY be omitted.
	ID string
	// ExtensionDefinitions is used to attach additional attributes and
	// associations to any BaseElement.
	// This association is not applicable when the XML schema interchange is used,
	// since the XSD mechanisms for supporting anyAttribute and any element already satisfy this requirement.
	ExtensionDefinitions []*ExtensionDefinition
	// Documentation is used to annotate the BPMN element, such as descriptions and other documentation.
	Documentation []*Documentation
	// ExtensionValues is used to provide values for extended attributes
	// and model associations.
	// This association is not applicable when the XML schema interchange is used,
	// since the XSD mechanisms for supporting anyAttribute and any element already satisfy this requirement.
	ExtensionValues []*ExtensionAttributeValue
}

// Documentation All BPMN elements that inherit from the BaseElement
// will have the capability, through the Documentation element,
// to have one (1) or more text descriptions of that element.
type Documentation struct {
	*BaseElement

	// Text is used to capture the text descriptions of a BPMN element.
	Text string

	// TextFormat identifies the format of the text.
	// It MUST follow the mime-type format. The default is "text/plain."
	TextFormat string
}

// RootElement is the abstract super class for all BPMN elements that are contained within Definitions
type RootElement struct {
	*BaseElement
	Definition *Definitions
}

type RelationshipDirection string

const (
	RelationshipDirectionNone     = RelationshipDirection("None")
	RelationshipDirectionForward  = RelationshipDirection("Forward")
	RelationshipDirectionBackward = RelationshipDirection("Backward")
	RelationshipDirectionBoth     = RelationshipDirection("Both")
)

// Extension element binds/imports an ExtensionDefinition and its attributes to a BPMN model definition.
type Extension struct {

	// MustUnderstand defines if the semantics defined by the extension
	// definition and its attribute definition MUST be understood by the BPMN adopter in order to process the BPMN model correctly.
	MustUnderstand bool

	// Definition of the content of the extension.
	// Note that in the XML schema, this definition is provided by an external XML schema file and is simply referenced by QName.
	Definition *ExtensionDefinition
}

// ExtensionDefinition defines and groups additional attributes.
type ExtensionDefinition struct {
	*Expression

	// Name of the extension. This is used as a namespace to uniquely identify the extension content.
	Name string

	ExtensionAttributeDefinitions []*ExtensionAttributeDefinition

	ExtensionValues []*ExtensionAttributeValue
}

// ExtensionAttributeValue contains the attribute value
type ExtensionAttributeValue struct {

	// Value the contained attribute value, used when the associated ExtensionAttributeDefinition.isReference is false.
	// The type of this Element MUST conform to the type specified in the associated ExtensionAttributeDefinition.
	Value *Element

	// ValueRef is the referenced attribute value,
	// used when the associated ExtensionAttributeDefinition.isReference is true.
	// The type of this Element MUST conform to the type specified in the associated ExtensionAttributeDefinition.
	ValueRef *Element

	// ExtensionAttributeDefinition extension attribute for which this value is being provided.
	ExtensionAttributeDefinition *ExtensionAttributeDefinition
}

// ExtensionAttributeDefinition defines new attributes
type ExtensionAttributeDefinition struct {
	// Name of the extension attribute.
	Name string
	// Type that is associated with the attribute.
	Type string
	// IsReference indicates if the attribute value will be referenced or contained.
	IsReference bool
}

// Relationship between elements
type Relationship struct {
	*BaseElement
	// Type descriptive name of the element.
	Type string
	// Direction of the relationship.
	Direction RelationshipDirection
	// Sources defines artifacts that are augmented by the relationship.
	Sources []*Element

	// Targets defines artifacts used to extend the semantics of the source element(s).
	Targets []*Element
}

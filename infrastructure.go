package core

// Definitions is the outermost containing object for all BPMN elements.
// It defines the scope of visibility and the namespace for all contained elements.
// The interchange of BPMN files will always be through one or more Definitions.
type Definitions struct {
	*BaseElement

	// Name of the Definition.
	Name string

	// TargetNamespace  identifies the namespace associated with
	// the Definition and follows the convention established by XML Schema.
	TargetNamespace string

	// ExpressionLanguage identifies the formal Expression language used in
	// Expressions within the elements of this Definition.
	// The Default is “http://www.w3.org/1999/XPath”.
	// This value MAY be overridden on each individual formal Expression.
	// The language MUST be specified in a URI format.
	ExpressionLanguage string

	// TypeLanguage identifies the type system used by the elements of this Definition.
	// Defaults to http://www.w3.org/2001/XMLSchema.
	// This value can be overridden on each individual ItemDefinition.
	// The language MUST be specified in a URI format.
	TypeLanguage string

	// Exporter identifies the tool that is exporting the bpmn model file.
	Exporter string

	//ExporterVersion identifies the version of the tool that is exporting the bpmn model file.
	ExporterVersion string

	// RootElements lists the root elements that are at the root of this Definitions.
	// These elements can be referenced within this Definitions and are visible to other Definitions.
	RootElements []*RootElement

	// Relationships enables the extension and integration of BPMN models into larger system/development Processes.
	Relationships []*Relationship

	// Extenstions identifies extensions beyond the attributes and model associations in the base BPMN specification
	Extenstions []*Extension

	// Imports is used to import externally defined elements and make them available for use by elements within this Definitions.
	Imports []*Import

	// Diagrams lists the BPMNDiagrams that are contained within this Definitions
	Diagrams []*BPMNDiagram
}

// Import is used when referencing external element,
// either BPMN elements contained in other BPMN Definitions or non-BPMN elements.
type Import struct {
	// ImportType Identifies the type of document being imported by providing
	// an absolute URI that identifies the encoding language used in the document.
	// The value of the importType attribute MUST be set to http://www.w3.org/2001/XMLSchema
	// when importing XML Schema 1.0 documents,
	// to http://www.w3.org/TR/wsdl20/ when importing WSDL 2.0 documents,
	// and http://www.omg.org/spec/BPMN/20100524/MODEL when importing BPMN 2.0 documents.
	// Other types of documents MAY be supported.
	// Importing Xml Schema 1.0, WSDL 2.0 and BPMN 2.0 types MUST be supported.
	ImportType string
	// Location identifies the location of the imported element.
	Location string
	// Namespace identifies the namespace of the imported element.
	Namespace  string
	Definition *Definitions
}

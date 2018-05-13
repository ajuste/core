package core

type Collaboration struct {
	Name         string
	IsClosed     bool
	Artifacts    []*Artifact
	Choreography *Choreography
	MessageFlows []*MessageFlow
}

type MessageFlow struct {
	Name          string
	Collaboration *Collaboration
	Message       *Message
}

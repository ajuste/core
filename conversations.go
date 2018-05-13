package core

type ConversationNode struct {
	Name            string
	Collaboration   *Collaboration
	CorrelationKeys []*CorrelationKey
	MessageFlows    []*MessageFlow
}

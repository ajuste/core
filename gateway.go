package core

type GatewayDirection int

const (
	GatewayDirectionUnspecified = GatewayDirection(0)
	GatewayDirectionConverging  = GatewayDirection(1)
	GatewayDirectionDiverging   = GatewayDirection(2)
	GatewayDirectionMixer       = GatewayDirection(3)
)

type EventBasedGatewayType int

const (
	EventBasedGatewayTypeParallel  = EventBasedGatewayType(0)
	EventBasedGatewayTypeExclusive = EventBasedGatewayType(1)
)

// Gateway is used to control how the Process flows (how Tokens flow)
// through Sequence Flows as they converge and diverge within a Process
type Gateway struct {
	*FlowNode
	GatewayDirection GatewayDirection
}

type ExclusiveGateway struct {
	*Gateway
	Default *SequenceFlow
}

type InclusveGateway struct {
	*Gateway
	Default *SequenceFlow
}

type ParallelGateway struct {
	*Gateway
	Default *SequenceFlow
}

type ComplexGateway struct {
	*Gateway
	Default              *SequenceFlow
	ActivationCondiction *Expression
}

type EventBasedGateway struct {
	Instantiate      bool
	EventGatewayType EventBasedGatewayType
}

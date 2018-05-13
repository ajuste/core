package core

type Operation struct {
	Name           string
	Implementation *Element
	OutMessage     *Message
	InMessage      *Message
}

type Interface struct {
	*RootElement
	Operations     []*Operation
	Implementation *Element
}

type EndPoint struct {
	*RootElement
}

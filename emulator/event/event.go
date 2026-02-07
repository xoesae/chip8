package event

type EventType int

const (
	DisplayClearEventType EventType = iota
	DisplayUpdatedEventType
)

type Event interface {
	Type() EventType
}

type DisplayClearEvent struct{}

func (DisplayClearEvent) Type() EventType { return DisplayClearEventType }

type DisplayUpdatedEvent struct {
	Pixels [32][64]bool
}

func (DisplayUpdatedEvent) Type() EventType { return DisplayUpdatedEventType }

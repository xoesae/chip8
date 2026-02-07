package event

type EventType int

const (
	DisplayClearEventType EventType = iota
	PixelUpdatedEventType
)

type Event interface {
	Type() EventType
}

type DisplayClearEvent struct{}

func (DisplayClearEvent) Type() EventType { return DisplayClearEventType }

type PixelUpdatedEvent struct {
	State bool
	X, Y  int
}

func (PixelUpdatedEvent) Type() EventType { return PixelUpdatedEventType }

package event

// Event is an interface what describe common event
type Event interface {
	Name() string
	Payload() map[string]interface{}
}

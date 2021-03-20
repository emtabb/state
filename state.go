package state

type State interface {
	Init() State
	Label([]string) State
	Field(string, interface{}) State
	Property(map[string] interface{}) State

	GetProperty() map[string] interface{}
	GetLabel() []string
	GetField(string) interface{}

	ToArray() []State
	ToString() string
	Drop(string) error
	Sum() float64
}

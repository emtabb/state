package state

type State interface {
	Init() State
	Label([]string) State
	Field(string, interface{}) State
	Property(map[string] interface{}) State

	GetProperty() map[string] interface{}
	GetLabel() []string
	GetField(string) interface{}

	ToStates() States
	ToString() string
	Drop(string) error
	Sum() float64
}

type States interface {
	State
	Generate() States
	ByStates([]State) States
	Of(...State) States
	Add(State)
	Get(int) State
	Set(int, State)
	Size() int
	Remove(int)
	Pop()
	IsEmpty() bool
	ToArray() []State
	Clear()
}

type List struct {
	States
	data []State
}

func (list *List) Generate() States {
	list.data = make([]State, 0)
	return list
}

func (list *List) ByStates(states []State) States {
	list.data = states[:]
	return list
}

func (list *List) Add(state State) {
	list.data = append(list.data, state)
}

func (list *List) Of(states ...State) States {
	list.data = states[:]
	return list
}

func (list *List) Get(index int) State {
	return list.data[index]
}

func (list *List) Set(index int, state State) {
	list.data[index] = state
}

func (list *List) Size() int {
	return len(list.data)
}

func (list *List) Remove(index int) {
	list.data = append(list.data[:index], list.data[index + 1:]...)
}

func (list *List) Pop() {
	list.data = list.data[:list.Size() - 1]
}

func (list *List) IsEmpty() bool {
	return len(list.data) == 0
}

func (list *List) ToArray() []State {
	return list.data
}

func (list *List) Clear() {
	list.data = []State{}
}
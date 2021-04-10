package state

type List struct {
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

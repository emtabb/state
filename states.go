package state

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
package state

type ObserverS struct {
	State
}

func (o ObserverS) Init() ObserverS {

	return o
}

func (o ObserverS) Label([]string) ObserverS {
	return o
}

func (o ObserverS) Field(string, interface{}) ObserverS {

	return o
}
func (o ObserverS) Property(map[string] interface{}) ObserverS {
	return o
}

func (o ObserverS) Watcher() interface{} {
	return nil
}

func (o ObserverS) GetProperty() map[string] interface{} {
	return nil
}

func (o ObserverS) GetLabel() []string {
	return nil
}

func (o ObserverS) GetField(string) interface{} {
	return nil
}

func (o ObserverS) ToStates() States {
	return nil
}

func (o ObserverS) ToString() string {
	return ""
}

func (o ObserverS) Drop(string) error {
	return nil
}
func (o ObserverS) Sum() float64 {
	return 0
}

